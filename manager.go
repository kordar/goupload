package goupload

import (
	"context"
	"fmt"
	"io"
)

type UploaderManager struct {
	container map[string]BucketUploader
}

func NewManager() *UploaderManager {
	return &UploaderManager{container: map[string]BucketUploader{}}
}

func NewManagerWithUploader(handlers ...BucketUploader) *UploaderManager {
	manager := &UploaderManager{container: map[string]BucketUploader{}}
	for _, handler := range handlers {
		manager.container[handler.BucketName()] = handler
	}
	return manager
}

func (mgr *UploaderManager) Add(uploader ...BucketUploader) {
	for _, handler := range uploader {
		mgr.container[handler.BucketName()] = handler
	}
}

func (mgr *UploaderManager) GetHandler(bucketName string) BucketUploader {
	return mgr.container[bucketName]
}

func (mgr *UploaderManager) Buckets(bucketName string, opt interface{}) []Bucket {
	handler := mgr.GetHandler(bucketName)
	if handler == nil {
		return []Bucket{}
	}
	ctx := context.Background()
	return handler.RemoteBuckets(ctx, opt)
}

func (mgr *UploaderManager) PutFromFile(bucketName string, name string, filePath string, opt interface{}) error {
	handler := mgr.GetHandler(bucketName)
	if handler == nil {
		return fmt.Errorf("[goupload-%s] no valid processor found", bucketName)
	}
	ctx := context.Background()
	return handler.PutFromFile(ctx, name, filePath, opt)
}

func (mgr *UploaderManager) Put(bucketName string, name string, fd io.Reader, opt interface{}) error {
	handler := mgr.GetHandler(bucketName)
	if handler == nil {
		return fmt.Errorf("[goupload-%s] no valid processor found", bucketName)
	}
	ctx := context.Background()
	return handler.Put(ctx, name, fd, opt)
}

func (mgr *UploaderManager) PutString(bucketName string, name string, content string, opt interface{}) error {
	handler := mgr.GetHandler(bucketName)
	if handler == nil {
		return fmt.Errorf("[goupload-%s] no valid processor found", bucketName)
	}
	ctx := context.Background()
	return handler.PutString(ctx, name, content, opt)
}

func (mgr *UploaderManager) List(bucketName string, path string, next string, limit int, opt interface{}) ([]BucketObject, string) {
	handler := mgr.GetHandler(bucketName)
	if handler == nil {
		return []BucketObject{}, ""
	}
	ctx := context.Background()
	return handler.List(ctx, path, next, limit, opt)
}

func (mgr *UploaderManager) Del(bucketName string, name string, opt interface{}) error {
	handler := mgr.GetHandler(bucketName)
	if handler == nil {
		return fmt.Errorf("[goupload-%s] no valid processor found", bucketName)
	}
	ctx := context.Background()
	return handler.Del(ctx, name, opt)
}

func (mgr *UploaderManager) DelAll(bucketName string, dir string) {
	handler := mgr.GetHandler(bucketName)
	if handler == nil {
		return
	}
	ctx := context.Background()
	mgr.GetHandler(bucketName).DelAll(ctx, dir)
}

func (mgr *UploaderManager) DelMulti(bucketName string, objects []BucketObject) error {
	handler := mgr.GetHandler(bucketName)
	if handler == nil {
		return fmt.Errorf("[goupload-%s] no valid processor found", bucketName)
	}
	ctx := context.Background()
	return handler.DelMulti(ctx, objects)
}

func (mgr *UploaderManager) Get(bucketName string, name string, opt interface{}) ([]byte, error) {
	handler := mgr.GetHandler(bucketName)
	if handler == nil {
		return nil, fmt.Errorf("[goupload-%s] no valid processor found", bucketName)
	}
	ctx := context.Background()
	return handler.Get(ctx, name, opt)
}

func (mgr *UploaderManager) GetToFile(bucketName string, name string, localPath string, opt interface{}) error {
	handler := mgr.GetHandler(bucketName)
	if handler == nil {
		return fmt.Errorf("[goupload-%s] no valid processor found", bucketName)
	}
	ctx := context.Background()
	return handler.GetToFile(ctx, name, localPath, opt)
}

func (mgr *UploaderManager) IsExist(bucketName string, name string) (bool, error) {
	handler := mgr.GetHandler(bucketName)
	if handler == nil {
		return false, fmt.Errorf("[goupload-%s] no valid processor found", bucketName)
	}
	ctx := context.Background()
	return handler.IsExist(ctx, name)
}

func (mgr *UploaderManager) Copy(bucketName string, dest string, source string, opt interface{}) error {
	handler := mgr.GetHandler(bucketName)
	if handler == nil {
		return fmt.Errorf("[goupload-%s] no valid processor found", bucketName)
	}
	ctx := context.Background()
	return handler.Copy(ctx, dest, source, opt)
}

func (mgr *UploaderManager) Move(bucketName string, dest string, source string, opt interface{}) error {
	handler := mgr.GetHandler(bucketName)
	if handler == nil {
		return fmt.Errorf("[goupload-%s] no valid processor found", bucketName)
	}
	ctx := context.Background()
	return handler.Move(ctx, dest, source, opt)
}

func (mgr *UploaderManager) Rename(bucketName string, dest string, source string, opt interface{}) error {
	handler := mgr.GetHandler(bucketName)
	if handler == nil {
		return fmt.Errorf("[goupload-%s] no valid processor found", bucketName)
	}
	ctx := context.Background()
	return handler.Rename(ctx, dest, source, opt)
}

func (mgr *UploaderManager) Tree(bucketName string, path string, next string, limit int, dep int, maxDep int, noleaf bool) []BucketTreeObject {
	handler := mgr.GetHandler(bucketName)
	if handler == nil {
		return []BucketTreeObject{}
	}
	ctx := context.Background()
	return handler.Tree(ctx, path, next, limit, dep, maxDep, noleaf)
}

func (mgr *UploaderManager) Append(bucketName string, name string, position int, r io.Reader, opt interface{}) (int, error) {
	handler := mgr.GetHandler(bucketName)
	if handler == nil {
		return -1, fmt.Errorf("[goupload-%s] no valid processor found", bucketName)
	}
	ctx := context.Background()
	return handler.Append(ctx, name, position, r, opt)
}

func (mgr *UploaderManager) AppendString(bucketName string, name string, position int, content string, opt interface{}) (int, error) {
	handler := mgr.GetHandler(bucketName)
	if handler == nil {
		return -1, fmt.Errorf("[goupload-%s] no valid processor found", bucketName)
	}
	ctx := context.Background()
	return handler.AppendString(ctx, name, position, content, opt)
}
