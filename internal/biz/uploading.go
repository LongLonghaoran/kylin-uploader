package biz

import "github.com/google/uuid"

type Uploading struct {
	Upid       string `json:"upid"`
	Filename   string `json:"filename"`
	TotalSize  int32  `json:"total_size"`
	MD5SUM     string `json:"md5sum"`
	Path       string `json:"path"`
	CurrentNum int32  `json:"current_num"`
	TotalCount int32  `json:"total_count"`
}

// Chunk is a Chunk model.
type Chunk struct {
	CID  int    `json:"cid"`
	Upid string `json:"upid"`
	Num  int32  `json:"num"`
	Size int32  `json:"size"`
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
