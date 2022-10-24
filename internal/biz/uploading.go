package biz

import "github.com/google/uuid"

type Uploading struct {
	Upid       string `json:"upid"`
	Filename   string `json:"filename"`
	TotalSize  int64  `json:"total_size"`
	MD5SUM     string `json:"md5sum"`
	Path       string `json:"path"`
	CurrentNum int64  `json:"current_num"`
	TotalCount int64  `json:"total_count"`
}

// Chunk is a Chunk model.
type Chunk struct {
	CID  int    `json:"cid"`
	Upid string `json:"upid"`
	Num  int64  `json:"num"`
	Size int64  `json:"size"`
	Path string `json:"path"`
}

func (u Uploading) ID() (jsonField string, value interface{}) {
	value = u.Upid
	jsonField = "upid"
	return
}

func (u Chunk) ID() (jsonField string, value interface{}) {
	value = uuid.NewString()
	jsonField = "cid"
	return
}
