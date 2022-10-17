package service

import (
	"context"
	"fmt"
	"path"

	pb "kylin-uploader/api/v1"
	"kylin-uploader/internal/biz"
)

type ChunkService struct {
	pb.UnimplementedChunkServer
	uc *biz.ChunkUsecase
}

func NewChunkService(uc *biz.ChunkUsecase) *ChunkService {
	return &ChunkService{uc: uc}
}

func (s *ChunkService) CreateUpload(ctx context.Context, req *pb.CreateUploadRequest) (*pb.CreateUploadReply, error) {
	savedUploading, err := s.uc.CreateUpload(req)
	if err != nil {
		return nil, fmt.Errorf("创建失败! %v", err)
	}
	return &pb.CreateUploadReply{UploadPath: "/api/v1/uploaders/" + savedUploading.Upid}, nil
}
func (s *ChunkService) UploadChunk(ctx context.Context, req *pb.UploadChunkRequest) (*pb.UploadChunkReply, error) {
	nextNum, err := s.uc.UploadChunk(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("上传失败!%v", err)
	}
	return &pb.UploadChunkReply{
		Nextnum: nextNum,
	}, nil
}
func (s *ChunkService) DoneUpload(ctx context.Context, req *pb.DoneUploadRequest) (*pb.DoneUploadReply, error) {
	uploading, err := s.uc.DoneUpload(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return &pb.DoneUploadReply{
		Path: path.Join("/files", uploading.Upid),
	}, nil
}
func (s *ChunkService) CheckFileExists(ctx context.Context, req *pb.CheckFileExistRequest) (*pb.CheckFileExistReply, error) {
	path, err := s.uc.CheckFileExists(req)
	if err != nil {
		return &pb.CheckFileExistReply{
			Exists: false,
			Path:   "",
		}, err
	} else {
		return &pb.CheckFileExistReply{
			Exists: true,
			Path:   path,
		}, nil
	}
}
func (s *ChunkService) CheckChunkExists(ctx context.Context, req *pb.CheckChunkExistsRequest) (*pb.CheckChunkExistsReply, error) {
	exists, err := s.uc.CheckChunkExists(req)
	return &pb.CheckChunkExistsReply{
		Exists: exists,
	}, err
}
