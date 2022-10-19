package data

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sync"

	"kylin-uploader/internal/biz"

	pb "kylin-uploader/api/v1"

	"github.com/go-kratos/kratos/v2/log"
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
	// 按照给定的filename,md5sum在所有json文件中查找指定的uploading是否存在
	entries, err := os.ReadDir(filepath.Join(chunkBasicDir, "index", "Uploading"))
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		up := biz.Uploading{}
		if entry.IsDir() {
			continue
		}
		f, err := os.Open(filepath.Join(chunkBasicDir, "index", "Uploading", entry.Name()))
		if err != nil {
			return nil, err
		}
		result, err := ioutil.ReadAll(f)
		if err != nil {
			return nil, err
		}
		array := make([]interface{}, 0)
		err = json.Unmarshal(result, &array)
		if len(array) == 0 || err != nil {
			continue
		}
		result, err = json.Marshal(array[0])
		if err != nil {
			return nil, err
		}
		json.Unmarshal(result, &up)
		if up.Filename == g.Filename && up.MD5SUM == g.MD5SUM {
			fmt.Println("uploading exists!", up.Filename, up.MD5SUM)
			return &up, nil
		}
	}

	// 不存在相同文件的情况
	err = r.data.DB.Insert(*g)
	if err != nil {
		return nil, err
	}
	err = os.Mkdir(path.Join(chunkBasicDir, g.Upid), 0766)
	if err != nil {
		return nil, err
	} else {
		return g, nil
	}
}

func (r *chunkRepo) FindChunk(req *pb.UploadChunkRequest) (*biz.Chunk, error) {
	var chunk biz.Chunk
	err := r.data.DB.Open(biz.Chunk{Upid: req.Upid}).Where("Num", "=", req.Num).First().AsEntity(&chunk)
	if err != nil {
		return nil, err
	}
	return &chunk, nil
}

func (r *chunkRepo) UploadChunk(req *pb.UploadChunkRequest, chunkBasicDir string) (*biz.Chunk, error) {
	var uploading = biz.Uploading{}
	err := r.data.DB.Open(biz.Uploading{Upid: req.Upid}).First().AsEntity(&uploading)
	if err != nil {
		log.Errorf("uploading找不到!%v", err)
		return nil, err
	}
	if req.Num == uploading.CurrentNum+1 {
		chunk := biz.Chunk{
			Upid: req.Upid,
			Num:  req.Num,
			Size: req.Size,
			Path: path.Join(chunkBasicDir, req.Upid, req.Upid+"."+fmt.Sprint(req.Num)),
		}
		r.data.DB.Insert(chunk)
		uploading.CurrentNum = uploading.CurrentNum + 1
		r.data.DB.Update(uploading)
		f, err := os.OpenFile(path.Join(chunkBasicDir, req.Upid, req.Upid+"."+fmt.Sprint(req.Num)), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			log.Errorf("failed to create new file!%v", err)
			return nil, err
		}
		// copy stream to file
		f.Write(req.Chunk)
		return &chunk, nil
	} else {
		chunk := biz.Chunk{}
		r.data.DB.Open(biz.Chunk{Upid: req.Upid}).Where("Num", "=", req.Num).First().AsEntity(&chunk)
		return &chunk, nil
	}
}

func (r *chunkRepo) DoneUpload(req *pb.DoneUploadRequest, chunkBasicDir string) (*biz.Uploading, error) {
	// 按upload_id查找分片信息:分片数
	var uploading biz.Uploading
	var chunks []biz.Chunk
	err := r.data.DB.Open(biz.Uploading{Upid: req.Upid}).First().AsEntity(&uploading)
	if err != nil {
		fmt.Println("-----------uploading not found")
		return nil, err
	}
	err = r.data.DB.Open(biz.Chunk{Upid: req.Upid}).Get().AsEntity(&chunks)
	if err != nil {
		fmt.Println("-----------chunk not found")
		return nil, err
	}
	if uploading.CurrentNum == uploading.TotalCount && len(uploading.Path) != 0 {
		return &uploading, nil
	}
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
	err = os.Rename(path.Join(chunkBasicDir, req.Upid, finalName), path.Join(chunkBasicDir, "files", uploading.Upid))
	if err != nil {
		return nil, err
	}
	uploading.Path = path.Join(chunkBasicDir, uploading.Upid)
	r.data.DB.Update(uploading)
	return &uploading, nil
}

func (r *chunkRepo) FindUploadingByUpid(upid string) (*biz.Uploading, error) {
	var uploading biz.Uploading
	err := r.data.DB.Open(biz.Uploading{Upid: upid}).First().AsEntity(&uploading)
	if err != nil {
		return nil, err
	}
	return &uploading, nil
}
func (r *chunkRepo) FindUploadingByFilename(filename, md5sum, chunkBasicDir string) (*biz.Uploading, error) {
	// 按照给定的filename,md5sum在所有json文件中查找指定的uploading是否存在
	entries, err := os.ReadDir(filepath.Join(chunkBasicDir, "index", "Uploading"))
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		up := biz.Uploading{}
		if entry.IsDir() {
			continue
		}
		f, err := os.Open(filepath.Join(chunkBasicDir, "index", "Uploading", entry.Name()))
		if err != nil {
			return nil, err
		}
		result, err := ioutil.ReadAll(f)
		if err != nil {
			return nil, err
		}
		array := make([]interface{}, 0)
		err = json.Unmarshal(result, &array)
		if len(array) == 0 || err != nil {
			continue
		}
		result, err = json.Marshal(array[0])
		if err != nil {
			return nil, err
		}
		json.Unmarshal(result, &up)
		if up.Filename == filename && up.MD5SUM == md5sum {
			fmt.Println("uploading exists!", up.Filename, up.MD5SUM)
			return &up, nil
		}
	}
	return nil, fmt.Errorf("file not exists")
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
