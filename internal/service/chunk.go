package service

import (
	"context"
	"fmt"

	pb "kylin-uploader/api/v1"
	"kylin-uploader/internal/biz"

	"github.com/google/uuid"
)

type ChunkService struct {
	pb.UnimplementedChunkServer
	uc *biz.ChunkUsecase
}

func NewChunkService(uc *biz.ChunkUsecase) *ChunkService {
	return &ChunkService{uc: uc}
}

func (s *ChunkService) CreateUpload(ctx context.Context, req *pb.CreateUploadRequest) (*pb.CreateUploadReply, error) {
	uploading := biz.Uploading{
		UUID:      uuid.NewString(),
		Filename:  req.Filename,
		TotalSize: req.TotalSize,
		MD5SUM:    req.Md5Sum,
	}
	savedUploading, err := s.uc.CreateUpload(context.Background(), &uploading)
	if err != nil {
		fmt.Println(fmt.Errorf("创建失败! %v", err))
	}
	return &pb.CreateUploadReply{UploadPath: "/api/v1/uploaders/" + savedUploading.UUID}, nil
}
func (s *ChunkService) UploadChunk(ctx context.Context, req *pb.UploadChunkRequest) (*pb.UploadChunkReply, error) {
	err := s.uc.UploadChunk(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("上传失败!%v", err)
	}
	return &pb.UploadChunkReply{}, nil
}
func (s *ChunkService) DoneUpload(ctx context.Context, req *pb.DoneUploadRequest) (*pb.DoneUploadReply, error) {
	return &pb.DoneUploadReply{}, nil
}
func (s *ChunkService) CheckFileExists(ctx context.Context, req *pb.CheckFileExistRequest) (*pb.CheckFileExistReply, error) {
	return &pb.CheckFileExistReply{}, nil
}
func (s *ChunkService) CheckChunkExists(ctx context.Context, req *pb.CheckChunkExistsRequest) (*pb.CheckChunkExistsReply, error) {
	return &pb.CheckChunkExistsReply{}, nil
}
