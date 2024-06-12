package goupload

import "io"

type IUpload interface {
	GetBucketName() string
	CreateBucket()
	Buckets() []Bucket
	PutFromFile(name string, filePath string) error
	Put(name string, fd io.Reader) error
	PutString(name string, content string) error
	List(path string, next string, limit int) BucketResult
	Del(name string) error
	Get(name string) ([]byte, error)
	GetToFile(name string, localPath string) error
}

type Bucket struct {
	Name       string `json:"name"`
	Region     string `json:"region"`
	CreateTime string `json:"create_time"`
}

type BucketResult struct {
	Content []Object `json:"content"`
	Dirs    []string `json:"dirs"`
}

type Object struct {
	Path      string `json:"path"`            // 存储路径
	Timestamp int64  `json:"timestamp"`       // 存储时间
	Size      int64  `json:"size"`            // 文件大小
	Type      string `json:"type,omitempty"`  // 文件类型
	Width     int    `json:"width,omitempty"` // 图片资源，额外宽高信息
	Height    int    `json:"height,omitempty"`
}
