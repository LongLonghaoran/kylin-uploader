package biz

import (
	"context"
	"path"

	pb "kylin-uploader/api/v1"
	v1 "kylin-uploader/api/v1"
	"kylin-uploader/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Chunk is a Chunk model.
type Chunk struct {
	gorm.Model
	UploadingID uint
	Num         int32
	Size        int32
	Path        string
	Uploading   Uploading
}

// ChunkRepo is a Greater repo.
type ChunkRepo interface {
	CreateUpload(*Uploading, string) (*Uploading, error)
	FindChunk(*pb.UploadChunkRequest) (*Chunk, *Uploading, error)
	UploadChunk(*pb.UploadChunkRequest, string) (*Chunk, error)
	DoneUpload(*pb.DoneUploadRequest, string) (*Uploading, error)
	FindUploader(string) (*Uploading, error)
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

func (uc *ChunkUsecase) UploadChunk(ctx context.Context, req *pb.UploadChunkRequest) (int32, error) {
	// find uploading
	chunk, uploading, err := uc.repo.FindChunk(req)
	if err == nil {
		log.Infof("chunk exists! %v", req.Upid)

		if chunk.Num < uploading.TotalCount {
			return chunk.Num + 1, nil
		} else {
			return -1, nil
		}
	}
	chunk, err = uc.repo.UploadChunk(req, chunkBasicDir)
	if err != nil {
		log.Errorf("Failed to Upload chunk!%v", err)
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
	uploading, err := uc.repo.FindUploader(req.Filename)
	if err != nil {
		return "", err
	}
	return path.Join("files", uploading.Filename), nil
}

func (uc *ChunkUsecase) CheckChunkExists(req *pb.CheckChunkExistsRequest) (bool, error) {
	_, _, err := uc.repo.FindChunk(&v1.UploadChunkRequest{Upid: req.Upid, Num: req.Num})
	if err != nil {
		return false, err
	}
	return true, nil
}
