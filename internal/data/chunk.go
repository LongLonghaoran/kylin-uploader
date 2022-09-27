package data

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"sync"

	"kylin-uploader/internal/biz"

	pb "kylin-uploader/api/v1"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type chunkRepo struct {
	data *Data
	log  *log.Helper
}

// NewchunkRepo .
func NewChunkRepo(data *Data, logger log.Logger) biz.ChunkRepo {
	return &chunkRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *chunkRepo) CreateUpload(g *biz.Uploading, chunkBasicDir string) (*biz.Uploading, error) {
	// 创建一条上传记录
	r.data.DB.Create(g)
	// create file & chunk record
	os.Mkdir(path.Join(chunkBasicDir, g.Upid), 0766)
	return g, nil
}

func (r *chunkRepo) FindChunk(req *pb.UploadChunkRequest) (*biz.Chunk, *biz.Uploading, error) {
	var uploading biz.Uploading
	r.data.DB.First(&uploading, "Upid = ?", req.Upid)
	var chunks []biz.Chunk
	r.data.DB.Model(&uploading).Where("Num = ?", req.Num).Association("Chunks").Find(&chunks)
	if len(chunks) == 0 {
		return nil, &uploading, fmt.Errorf("chunk not exists %v", req.Upid)
	} else {
		return &chunks[0], &uploading, nil
	}
}

func (r *chunkRepo) UploadChunk(req *pb.UploadChunkRequest, chunkBasicDir string) (*biz.Chunk, error) {
	var uploading biz.Uploading
	r.data.DB.First(&uploading, "Upid = ?", req.Upid)
	chunk := biz.Chunk{
		UploadingID: uint(uploading.ID),
		Num:         req.Num,
		Size:        req.Size,
	}
	err := r.data.DB.Transaction(func(tx *gorm.DB) error {
		// create chunk record
		result := r.data.DB.Create(&chunk)
		if result.Error != nil {
			return result.Error
		}
		f, err := os.OpenFile(path.Join(chunkBasicDir, req.Upid, req.Upid+"."+fmt.Sprint(req.Num)), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			log.Errorf("failed to create new file!%v", err)
			return err
		}
		// copy stream to file
		f.WriteString(req.Chunk)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &chunk, nil
}

func (r *chunkRepo) DoneUpload(req *pb.DoneUploadRequest, chunkBasicDir string) (*biz.Uploading, error) {
	// 按upload_id查找分片信息:分片数
	var uploading biz.Uploading
	var chunks []biz.Chunk
	r.data.DB.First(&uploading, "upid = ?", req.Upid)
	r.data.DB.Model(&uploading).Association("Chunks").Find(&chunks)
	// 核对分片数
	if int(uploading.TotalCount) != len(chunks) {
		return nil, fmt.Errorf("分片总数错误, %v %v", uploading.TotalCount, len(chunks))
	}
	fis, err := ioutil.ReadDir(path.Join(chunkBasicDir, req.Upid))
	if err != nil {
		return nil, err
	}
	chunkFileNames := make([]string, 0)
	for _, fi := range fis {
		if !fi.IsDir() {
			chunkFileNames = append(chunkFileNames, fi.Name())
		}
	}
	finalName, _ := RecursiveMergeChunk(path.Join(chunkBasicDir, req.Upid), chunkFileNames...)
	fmt.Println("finalName-----------------", finalName)
	err = os.Rename(path.Join(chunkBasicDir, req.Upid, finalName), path.Join(chunkBasicDir, uploading.Filename))
	if err != nil {
		return nil, err
	}
	uploading = biz.Uploading{
		Path: path.Join(chunkBasicDir, uploading.Filename),
	}
	r.data.DB.Model(&biz.Uploading{}).Where("upid = ?", req.Upid).Update("path", path.Join(chunkBasicDir, req.Upid, uploading.Filename))
	return &uploading, nil
}

func RecursiveMergeChunk(chunkBasicDir string, chunkFileNames ...string) (finalName string, e error) {
	// var maxMem int64 = 1 * 1024 * 1024 * 1024 // max memery: 1GB
	if len(chunkFileNames) == 1 {
		return chunkFileNames[0], nil
	}
	var perTimes, goRoutineCount int
	totalCount := len(chunkFileNames)
	if totalCount < 5 {
		perTimes = totalCount
	} else {
		perTimes = 5
	}
	goRoutineCount = len(chunkFileNames) / perTimes
	if goRoutineCount*perTimes < len(chunkFileNames) {
		goRoutineCount += 1
	}
	wg := &sync.WaitGroup{}
	errChan := make(chan bool, goRoutineCount)
	newChunkFileNames := make([]string, 0)
	for i := 0; i < goRoutineCount; i++ {
		newChunkFileNames = append(newChunkFileNames, chunkFileNames[i*int(perTimes)])
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			fw, err := os.OpenFile(
				path.Join(chunkBasicDir, chunkFileNames[j*int(perTimes)]), os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				log.Error("os.OpenFile error: %v", err)
				errChan <- true
				return
			}
			defer fw.Close()
			for k := j * int(perTimes); k < int(perTimes)*(j+1); k++ {
				if k > int(len(chunkFileNames))-1 {
					break
				}

				// Skip the first file.
				if k == j*int(perTimes) {
					continue
				}

				file, err := os.Open(path.Join(chunkBasicDir, chunkFileNames[k]))
				if err != nil {
					log.Error("os.Open error: %v", err)
					errChan <- true
					return
				}
				_, err = io.Copy(fw, file)
				if err != nil {
					log.Error("io.Copy error: %v", err)
					_ = file.Close()

					errChan <- true
					return
				}
				err = fw.Sync()
				if err != nil {
					log.Error("fw.Sync error: %v", err)
					errChan <- true
					return
				}
				_ = file.Close()
				err = os.RemoveAll(path.Join(chunkBasicDir, chunkFileNames[k]))
				if err != nil {
					log.Error("os.RemoveAll error: %v", err)
					errChan <- true
					return
				}
			}
			errChan <- false
		}(i)
	}

	wg.Wait()
	for i := 0; i < len(errChan); i++ {
		if <-errChan {
			return "", fmt.Errorf("merge file error")
		}
	}

	finalName, err := RecursiveMergeChunk(chunkBasicDir, newChunkFileNames...)
	if err != nil {
		log.Error("mergeChunkRecursive error: %v", err)
		return "", err
	}
	return finalName, nil
}
