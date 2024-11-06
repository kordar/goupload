package goupload

import (
	"context"
	"io"
)

type BucketUploader interface {
	BucketName() string
	DriverName() string
	RemoteBuckets(ctx context.Context, opt interface{}) []Bucket
	Get(ctx context.Context, name string, opt interface{}, id ...string) ([]byte, error)
	GetToFile(ctx context.Context, name string, localPath string, opt interface{}, id ...string) error
	PutFromFile(ctx context.Context, name string, filePath string, opt interface{}) error
	Put(ctx context.Context, name string, fd io.Reader, opt interface{}) error
	PutString(ctx context.Context, name string, content string, opt interface{}) error
	List(ctx context.Context, path string, next string, limit int, opt interface{}) ([]BucketObject, string)
	Del(ctx context.Context, name string, opt interface{}) error
	DelAll(ctx context.Context, dir string)
	DelMulti(ctx context.Context, objects []BucketObject) error
	IsExist(ctx context.Context, name string, id ...string) (bool, error)
	Copy(ctx context.Context, dest string, source string, opt interface{}) error
	Move(ctx context.Context, dest string, source string, opt interface{}) error
	Rename(ctx context.Context, dest string, source string, opt interface{}) error
	Tree(ctx context.Context, path string, next string, limit int, dep int, maxDep int, noleaf bool) []BucketTreeObject
	Append(ctx context.Context, name string, position int, r io.Reader, opt interface{}) (int, error)
	AppendString(ctx context.Context, name string, position int, content string, opt interface{}) (int, error)
}
