package biz

import (
	"context"
	"fmt"
	"os"

	pb "kylin-uploader/api/v1"

	"github.com/go-kratos/kratos/v2/log"
)

// Chunk is a Chunk model.
type Chunk struct {
	Upid string
	Num  int32
	Size int32
	Path string
}

// ChunkRepo is a Greater repo.
type ChunkRepo interface {
	CreateUpload(context.Context, *Uploading) (*Uploading, error)
	FindChunk(context.Context, *pb.UploadChunkRequest) (*Chunk, error)
}

// ChunkUsecase is a Chunk usecase.
type ChunkUsecase struct {
	repo ChunkRepo
	log  *log.Helper
}

// NewChunkUsecase new a Chunk usecase.
func NewChunkUsecase(repo ChunkRepo, logger log.Logger) *ChunkUsecase {
	return &ChunkUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateChunk creates a Chunk, and returns the new Chunk.
func (uc *ChunkUsecase) CreateUpload(ctx context.Context, g *Uploading) (*Uploading, error) {
	log.Infof("CreateUpload: ", g)
	return uc.repo.CreateUpload(ctx, g)
}

func (uc *ChunkUsecase) UploadChunk(ctx context.Context, req *pb.UploadChunkRequest) error {
	log.Infof("UploadChunk: upid: %s, num: %s", req.Upid, req.Num)
	// find uploading
	_, err := uc.repo.FindChunk(ctx, req)
	if err == nil {
		log.Infof("File exists! %v", req.Upid)
		return nil
	}
	// create file
	f, err := os.OpenFile("/tmp/chunkdir/"+req.Upid+"."+fmt.Sprint(req.Num), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Errorf("Failed to Create new file!%v", err)
		return err
	}
	// copy stream to file
	f.WriteString(req.Chunk)
	return nil
}
