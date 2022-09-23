package data

import (
	"context"
	"fmt"
	"os"

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

func (r *chunkRepo) CreateUpload(ctx context.Context, g *biz.Uploading) (*biz.Uploading, error) {
	// 创建一条上传记录
	log.Info("internal>data>chunk#CreateUpload: ", g)
	return g, nil
}

func (r *chunkRepo) FindChunk(ctx context.Context, req *pb.UploadChunkRequest) (*biz.Chunk, error) {
	fi, err := os.Stat("/tmp/chunkdir/" + req.Upid + "." + fmt.Sprint(req.Num))
	if !os.IsNotExist(err) {
		return &biz.Chunk{
			Upid: req.Upid,
			Size: int32(fi.Size()),
			Path: "/tmp/chunkdir/" + fi.Name(),
		}, nil
	}
	return nil, fmt.Errorf("文件不存在 %s", "/tmp/chunkdir/"+req.Upid+"."+fmt.Sprint(req.Num))
}
