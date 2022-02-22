package goupload

import (
	"io"
)

type UploadManager struct {
	handler IUpload
}

func NewUploadManager(handler IUpload) *UploadManager {
	handler.CreateBucket()
	return &UploadManager{handler: handler}
}

func (mgr *UploadManager) Buckets() []Bucket {
	return mgr.handler.Buckets()
}

func (mgr *UploadManager) PutFromFile(name string, filePath string) error {
	return mgr.handler.PutFromFile(name, filePath)
}

func (mgr *UploadManager) Put(name string, fd io.Reader) error {
	return mgr.handler.Put(name, fd)
}

func (mgr *UploadManager) PutString(name string, content string) error {
	return mgr.handler.PutString(name, content)
}

func (mgr *UploadManager) List(path string, next string, limit int) BucketResult {
	return mgr.handler.List(path, next, limit)
}

func (mgr *UploadManager) Del(name string) error {
	return mgr.handler.Del(name)
}

func (mgr *UploadManager) Get(name string) ([]byte, error) {
	return mgr.handler.Get(name)
}

func (mgr *UploadManager) GetToFile(name string, localPath string) error {
	return mgr.handler.GetToFile(name, localPath)
}
