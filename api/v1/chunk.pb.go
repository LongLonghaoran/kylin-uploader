// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: api/v1/chunk.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 文件名
	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	// 总分片数
	TotalCount int32 `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	// 总容量
	TotalSize int32 `protobuf:"varint,3,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	// md5校验和
	Md5Sum string `protobuf:"bytes,4,opt,name=md5sum,proto3" json:"md5sum,omitempty"`
}

func (x *CreateUploadRequest) Reset() {
	*x = CreateUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_chunk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUploadRequest) ProtoMessage() {}

func (x *CreateUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_chunk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUploadRequest.ProtoReflect.Descriptor instead.
func (*CreateUploadRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_chunk_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUploadRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *CreateUploadRequest) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *CreateUploadRequest) GetTotalSize() int32 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

func (x *CreateUploadRequest) GetMd5Sum() string {
	if x != nil {
		return x.Md5Sum
	}
	return ""
}

type CreateUploadReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 之后用于上传分片的路径
	UploadPath string `protobuf:"bytes,1,opt,name=upload_path,json=uploadPath,proto3" json:"upload_path,omitempty"`
}

func (x *CreateUploadReply) Reset() {
	*x = CreateUploadReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_chunk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUploadReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUploadReply) ProtoMessage() {}

func (x *CreateUploadReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_chunk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUploadReply.ProtoReflect.Descriptor instead.
func (*CreateUploadReply) Descriptor() ([]byte, []int) {
	return file_api_v1_chunk_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUploadReply) GetUploadPath() string {
	if x != nil {
		return x.UploadPath
	}
	return ""
}

type UploadChunkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//上传的upid
	Upid string `protobuf:"bytes,1,opt,name=upid,proto3" json:"upid,omitempty"`
	//当前块的序号
	Num int32 `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	//当前块的大小
	Size int32 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	//当前块的数据
	Chunk string `protobuf:"bytes,4,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *UploadChunkRequest) Reset() {
	*x = UploadChunkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_chunk_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadChunkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadChunkRequest) ProtoMessage() {}

func (x *UploadChunkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_chunk_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadChunkRequest.ProtoReflect.Descriptor instead.
func (*UploadChunkRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_chunk_proto_rawDescGZIP(), []int{2}
}

func (x *UploadChunkRequest) GetUpid() string {
	if x != nil {
		return x.Upid
	}
	return ""
}

func (x *UploadChunkRequest) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *UploadChunkRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *UploadChunkRequest) GetChunk() string {
	if x != nil {
		return x.Chunk
	}
	return ""
}

type UploadChunkReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nextnum int32 `protobuf:"varint,2,opt,name=nextnum,proto3" json:"nextnum,omitempty"`
}

func (x *UploadChunkReply) Reset() {
	*x = UploadChunkReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_chunk_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadChunkReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadChunkReply) ProtoMessage() {}

func (x *UploadChunkReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_chunk_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadChunkReply.ProtoReflect.Descriptor instead.
func (*UploadChunkReply) Descriptor() ([]byte, []int) {
	return file_api_v1_chunk_proto_rawDescGZIP(), []int{3}
}

func (x *UploadChunkReply) GetNextnum() int32 {
	if x != nil {
		return x.Nextnum
	}
	return 0
}

type DoneUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 上传id
	Upid string `protobuf:"bytes,1,opt,name=upid,proto3" json:"upid,omitempty"`
}

func (x *DoneUploadRequest) Reset() {
	*x = DoneUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_chunk_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoneUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoneUploadRequest) ProtoMessage() {}

func (x *DoneUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_chunk_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoneUploadRequest.ProtoReflect.Descriptor instead.
func (*DoneUploadRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_chunk_proto_rawDescGZIP(), []int{4}
}

func (x *DoneUploadRequest) GetUpid() string {
	if x != nil {
		return x.Upid
	}
	return ""
}

type DoneUploadReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 返回的文件路径
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *DoneUploadReply) Reset() {
	*x = DoneUploadReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_chunk_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoneUploadReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoneUploadReply) ProtoMessage() {}

func (x *DoneUploadReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_chunk_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoneUploadReply.ProtoReflect.Descriptor instead.
func (*DoneUploadReply) Descriptor() ([]byte, []int) {
	return file_api_v1_chunk_proto_rawDescGZIP(), []int{5}
}

func (x *DoneUploadReply) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type CheckFileExistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 文件名
	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *CheckFileExistRequest) Reset() {
	*x = CheckFileExistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_chunk_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckFileExistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckFileExistRequest) ProtoMessage() {}

func (x *CheckFileExistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_chunk_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckFileExistRequest.ProtoReflect.Descriptor instead.
func (*CheckFileExistRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_chunk_proto_rawDescGZIP(), []int{6}
}

func (x *CheckFileExistRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type CheckFileExistReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 是否存在
	Exists bool `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
	// 存在则返回路径
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *CheckFileExistReply) Reset() {
	*x = CheckFileExistReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_chunk_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckFileExistReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckFileExistReply) ProtoMessage() {}

func (x *CheckFileExistReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_chunk_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckFileExistReply.ProtoReflect.Descriptor instead.
func (*CheckFileExistReply) Descriptor() ([]byte, []int) {
	return file_api_v1_chunk_proto_rawDescGZIP(), []int{7}
}

func (x *CheckFileExistReply) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

func (x *CheckFileExistReply) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type CheckChunkExistsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 上传id
	Upid string `protobuf:"bytes,1,opt,name=upid,proto3" json:"upid,omitempty"`
	// 分片序号
	Num int32 `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *CheckChunkExistsRequest) Reset() {
	*x = CheckChunkExistsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_chunk_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckChunkExistsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckChunkExistsRequest) ProtoMessage() {}

func (x *CheckChunkExistsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_chunk_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckChunkExistsRequest.ProtoReflect.Descriptor instead.
func (*CheckChunkExistsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_chunk_proto_rawDescGZIP(), []int{8}
}

func (x *CheckChunkExistsRequest) GetUpid() string {
	if x != nil {
		return x.Upid
	}
	return ""
}

func (x *CheckChunkExistsRequest) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type CheckChunkExistsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
}

func (x *CheckChunkExistsReply) Reset() {
	*x = CheckChunkExistsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_chunk_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckChunkExistsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckChunkExistsReply) ProtoMessage() {}

func (x *CheckChunkExistsReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_chunk_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckChunkExistsReply.ProtoReflect.Descriptor instead.
func (*CheckChunkExistsReply) Descriptor() ([]byte, []int) {
	return file_api_v1_chunk_proto_rawDescGZIP(), []int{9}
}

func (x *CheckChunkExistsReply) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

var File_api_v1_chunk_proto protoreflect.FileDescriptor

var file_api_v1_chunk_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x64, 0x35, 0x73, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6d, 0x64, 0x35, 0x73, 0x75, 0x6d, 0x22, 0x34, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x61, 0x74, 0x68, 0x22, 0x64, 0x0a, 0x12,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x70, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x22, 0x2c, 0x0a, 0x10, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x78, 0x74, 0x6e, 0x75,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6e, 0x65, 0x78, 0x74, 0x6e, 0x75, 0x6d,
	0x22, 0x27, 0x0a, 0x11, 0x44, 0x6f, 0x6e, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x70, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x0f, 0x44, 0x6f, 0x6e,
	0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x22, 0x33, 0x0a, 0x15, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x46, 0x69, 0x6c, 0x65, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x41, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x46, 0x69,
	0x6c, 0x65, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x78,
	0x69, 0x73, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x3f, 0x0a, 0x17, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x70, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x2f, 0x0a, 0x15, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x32, 0xaa, 0x04, 0x0a, 0x05, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x12, 0x64, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x68, 0x0a, 0x0b, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x1a, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x70, 0x69, 0x64,
	0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x67, 0x0a, 0x0a, 0x44, 0x6f, 0x6e, 0x65, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x6e, 0x65,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x6e, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x2f, 0x7b, 0x75, 0x70, 0x69, 0x64, 0x7d, 0x2f, 0x64, 0x6f, 0x6e, 0x65, 0x12, 0x70, 0x0a,
	0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x46, 0x69, 0x6c, 0x65, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73,
	0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x46,
	0x69, 0x6c, 0x65, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x46, 0x69,
	0x6c, 0x65, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12,
	0x76, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x42, 0x24, 0x0a, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x50, 0x01, 0x5a, 0x18, 0x6b, 0x79, 0x6c, 0x69, 0x6e, 0x2d, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_chunk_proto_rawDescOnce sync.Once
	file_api_v1_chunk_proto_rawDescData = file_api_v1_chunk_proto_rawDesc
)

func file_api_v1_chunk_proto_rawDescGZIP() []byte {
	file_api_v1_chunk_proto_rawDescOnce.Do(func() {
		file_api_v1_chunk_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_chunk_proto_rawDescData)
	})
	return file_api_v1_chunk_proto_rawDescData
}

var file_api_v1_chunk_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_v1_chunk_proto_goTypes = []interface{}{
	(*CreateUploadRequest)(nil),     // 0: api.v1.CreateUploadRequest
	(*CreateUploadReply)(nil),       // 1: api.v1.CreateUploadReply
	(*UploadChunkRequest)(nil),      // 2: api.v1.UploadChunkRequest
	(*UploadChunkReply)(nil),        // 3: api.v1.UploadChunkReply
	(*DoneUploadRequest)(nil),       // 4: api.v1.DoneUploadRequest
	(*DoneUploadReply)(nil),         // 5: api.v1.DoneUploadReply
	(*CheckFileExistRequest)(nil),   // 6: api.v1.CheckFileExistRequest
	(*CheckFileExistReply)(nil),     // 7: api.v1.CheckFileExistReply
	(*CheckChunkExistsRequest)(nil), // 8: api.v1.CheckChunkExistsRequest
	(*CheckChunkExistsReply)(nil),   // 9: api.v1.CheckChunkExistsReply
}
var file_api_v1_chunk_proto_depIdxs = []int32{
	0, // 0: api.v1.Chunk.CreateUpload:input_type -> api.v1.CreateUploadRequest
	2, // 1: api.v1.Chunk.UploadChunk:input_type -> api.v1.UploadChunkRequest
	4, // 2: api.v1.Chunk.DoneUpload:input_type -> api.v1.DoneUploadRequest
	6, // 3: api.v1.Chunk.CheckFileExists:input_type -> api.v1.CheckFileExistRequest
	8, // 4: api.v1.Chunk.CheckChunkExists:input_type -> api.v1.CheckChunkExistsRequest
	1, // 5: api.v1.Chunk.CreateUpload:output_type -> api.v1.CreateUploadReply
	3, // 6: api.v1.Chunk.UploadChunk:output_type -> api.v1.UploadChunkReply
	5, // 7: api.v1.Chunk.DoneUpload:output_type -> api.v1.DoneUploadReply
	7, // 8: api.v1.Chunk.CheckFileExists:output_type -> api.v1.CheckFileExistReply
	9, // 9: api.v1.Chunk.CheckChunkExists:output_type -> api.v1.CheckChunkExistsReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1_chunk_proto_init() }
func file_api_v1_chunk_proto_init() {
	if File_api_v1_chunk_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_chunk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUploadRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_chunk_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUploadReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_chunk_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadChunkRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_chunk_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadChunkReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_chunk_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoneUploadRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_chunk_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoneUploadReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_chunk_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckFileExistRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_chunk_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckFileExistReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_chunk_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckChunkExistsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_chunk_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckChunkExistsReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_chunk_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_chunk_proto_goTypes,
		DependencyIndexes: file_api_v1_chunk_proto_depIdxs,
		MessageInfos:      file_api_v1_chunk_proto_msgTypes,
	}.Build()
	File_api_v1_chunk_proto = out.File
	file_api_v1_chunk_proto_rawDesc = nil
	file_api_v1_chunk_proto_goTypes = nil
	file_api_v1_chunk_proto_depIdxs = nil
}
