package biz

import (
	"context"
	"fmt"
	"path"
	"strconv"

	pb "kylin-uploader/api/v1"
	"kylin-uploader/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/uuid"
)

// ChunkRepo is a Greater repo.
type ChunkRepo interface {
	CreateUpload(*Uploading, string) (*Uploading, error)
	FindChunk(*pb.UploadChunkRequest) (*Chunk, error)
	UploadChunk(http.Context, *pb.UploadChunkRequest, string) (*Chunk, error)
	DoneUpload(*pb.DoneUploadRequest, string) (*Uploading, error)
	FindUploadingByUpid(upid string) (*Uploading, error)
	FindUploadingByFilename(filename, md5sum, chunkBasicDir string) (*Uploading, error)
}

// ChunkUsecase is a Chunk usecase.
type ChunkUsecase struct {
	repo ChunkRepo
	log  *log.Helper
}

var chunkBasicDir string

// NewChunkUsecase new a Chunk usecase.
func NewChunkUsecase(repo ChunkRepo, s *conf.Server, logger log.Logger) *ChunkUsecase {
	chunkBasicDir = s.Basicdir
	return &ChunkUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateChunk creates a Chunk, and returns the new Chunk.
func (uc *ChunkUsecase) CreateUpload(req *pb.CreateUploadRequest) (*Uploading, error) {
	uploading := Uploading{
		Upid:       uuid.NewString(),
		Filename:   req.Filename,
		TotalSize:  req.TotalSize,
		TotalCount: req.TotalCount,
		MD5SUM:     req.Md5Sum,
	}

	return uc.repo.CreateUpload(&uploading, chunkBasicDir)
}

func (uc *ChunkUsecase) UploadChunk(ctx http.Context, req *pb.UploadChunkRequest) (int64, error) {
	// find uploading
	num, err := strconv.Atoi(ctx.Query()["num"][0])
	if err != nil {
		fmt.Println("转换参数错误", err)
		return -1, err
	}
	size, err := strconv.Atoi(ctx.Query()["size"][0])
	if err != nil {
		fmt.Println("转换参数错误", err)
		return -1, err
	}
	req.Num = int64(num)
	req.Size = int64(size)
	uploading, _ := uc.repo.FindUploadingByUpid(req.Upid)
	if uploading == nil {
		return -1, fmt.Errorf("uploading不存在")
	}
	chunk, err := uc.repo.FindChunk(req)
	if err == nil {
		log.Infof("chunk exists! %v", req.Upid)

		if chunk.Num < uploading.TotalCount {
			return chunk.Num + 1, nil
		} else {
			return -1, nil
		}
	}
	chunk, err = uc.repo.UploadChunk(ctx, req, chunkBasicDir)
	if err != nil {
		return 0, err
	}
	if chunk.Num < uploading.TotalCount {
		return chunk.Num + 1, nil
	} else {
		return -1, nil
	}
}
func (uc *ChunkUsecase) DoneUpload(ctx context.Context, req *pb.DoneUploadRequest) (*Uploading, error) {
	uploading, err := uc.repo.DoneUpload(req, chunkBasicDir)
	if err != nil {
		return nil, err
	}
	return uploading, nil
}

func (uc *ChunkUsecase) CheckFileExists(req *pb.CheckFileExistRequest) (string, error) {
	uploading, err := uc.repo.FindUploadingByFilename(req.Filename, req.Md5Sum, chunkBasicDir)
	if err != nil {
		return "", err
	}
	return path.Join("files", uploading.Upid), nil
}

func (uc *ChunkUsecase) CheckChunkExists(req *pb.CheckChunkExistsRequest) (bool, error) {
	_, err := uc.repo.FindChunk(&pb.UploadChunkRequest{Upid: req.Upid, Num: req.Num})
	if err != nil {
		return false, err
	}
	return true, nil
}
