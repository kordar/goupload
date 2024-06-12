package goupload

import (
	"io"
)

type BucketManager struct {
	bucketHandler map[string]IUpload
}

func NewBucketManager() *BucketManager {
	return &BucketManager{bucketHandler: map[string]IUpload{}}
}

func NewBucketManagers(handlers []IUpload, autoCreateBucket bool) *BucketManager {
	manager := &BucketManager{bucketHandler: map[string]IUpload{}}
	for _, handler := range handlers {
		if autoCreateBucket {
			handler.CreateBucket()
		}
		manager.bucketHandler[handler.GetBucketName()] = handler
	}
	return manager
}

func (mgr *BucketManager) SetUploadHandlers(handler ...IUpload) {
	for _, upload := range handler {
		mgr.bucketHandler[upload.GetBucketName()] = upload
	}
}

func (mgr *BucketManager) GetHandler(bucketName string) IUpload {
	return mgr.bucketHandler[bucketName]
}

func (mgr *BucketManager) CreateBucket(bucketName string) {
	mgr.bucketHandler[bucketName].CreateBucket()
}

func (mgr *BucketManager) Buckets(bucketName string) []Bucket {
	return mgr.GetHandler(bucketName).Buckets()
}

func (mgr *BucketManager) PutFromFile(bucketName string, name string, filePath string) error {
	return mgr.GetHandler(bucketName).PutFromFile(name, filePath)
}

func (mgr *BucketManager) Put(bucketName string, name string, fd io.Reader) error {
	return mgr.GetHandler(bucketName).Put(name, fd)
}

func (mgr *BucketManager) PutString(bucketName string, name string, content string) error {
	return mgr.GetHandler(bucketName).PutString(name, content)
}

func (mgr *BucketManager) List(bucketName string, path string, next string, limit int) BucketResult {
	return mgr.GetHandler(bucketName).List(path, next, limit)
}

func (mgr *BucketManager) Del(bucketName string, name string) error {
	return mgr.GetHandler(bucketName).Del(name)
}

func (mgr *BucketManager) Get(bucketName string, name string) ([]byte, error) {
	return mgr.GetHandler(bucketName).Get(name)
}

func (mgr *BucketManager) GetToFile(bucketName string, name string, localPath string) error {
	return mgr.GetHandler(bucketName).GetToFile(name, localPath)
}
