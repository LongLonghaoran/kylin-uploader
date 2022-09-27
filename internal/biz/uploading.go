package biz

import "gorm.io/gorm"

type Uploading struct {
	gorm.Model
	Upid       string
	Filename   string
	TotalSize  int32
	MD5SUM     string
	Path       string
	CurrentNum int32
	TotalCount int32
	Chunks     []Chunk
}
