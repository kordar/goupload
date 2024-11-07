package goupload

import (
	"context"
	"io"
)

// BucketUploader defines a set of methods for uploading, downloading, managing, and manipulating objects in a storage bucket system.
type BucketUploader interface {
	// Name returns the name of the uploader.
	Name() string

	// Driver returns the driver or type of storage used (e.g., S3, local storage).
	Driver() string

	// RemoteBuckets lists available remote buckets based on provided context and arguments.
	RemoteBuckets(ctx context.Context, args ...interface{}) []Bucket

	// Get retrieves the content of an object by name and returns it as a byte slice.
	Get(ctx context.Context, name string, args ...interface{}) ([]byte, error)

	// GetToFile downloads an object by name and saves it to the specified local file path.
	GetToFile(ctx context.Context, name string, localPath string, args ...interface{}) error

	// Put uploads content from an io.Reader to the storage under the specified name.
	Put(ctx context.Context, name string, fd io.Reader, args ...interface{}) error

	// PutString uploads content from a string to the storage under the specified name.
	PutString(ctx context.Context, name string, content string, args ...interface{}) error

	// PutFromFile uploads content from a local file to the storage under the specified name.
	PutFromFile(ctx context.Context, name string, filePath string, args ...interface{}) error

	// List retrieves objects within a specified path, with options for pagination (next and limit).
	List(ctx context.Context, dir string, next interface{}, limit int, args ...interface{}) ([]BucketObject, interface{})

	// Del deletes an object by name from the storage.
	Del(ctx context.Context, name string, args ...interface{}) error

	// DelAll deletes all objects within a specified directory in the storage.
	DelAll(ctx context.Context, dir string, args ...interface{})

	// DelMulti deletes multiple objects specified in the objects slice from the storage.
	DelMulti(ctx context.Context, objects []BucketObject, args ...interface{}) error

	// IsExist checks if an object by name exists in the storage.
	IsExist(ctx context.Context, name string, args ...interface{}) (bool, error)

	// Copy duplicates an object from the source path to the destination path.
	Copy(ctx context.Context, dest string, source string, args ...interface{}) error

	// Move transfers an object from the source path to the destination path.
	Move(ctx context.Context, dest string, source string, args ...interface{}) error

	// Rename renames an object from the source name to the destination name.
	Rename(ctx context.Context, dest string, source string, args ...interface{}) error

	// Tree lists the structure of objects within a directory up to a specified depth, optionally excluding leaves.
	Tree(ctx context.Context, dir string, next interface{}, limit int, dep int, maxDep int, noleaf bool, args ...interface{}) []BucketTreeObject

	// Append appends data from an io.Reader to an existing object starting at the specified position.
	Append(ctx context.Context, name string, position int, fd io.Reader, args ...interface{}) (int, error)

	// AppendString appends data from a string to an existing object starting at the specified position.
	AppendString(ctx context.Context, name string, position int, content string, args ...interface{}) (int, error)
}

type BucketZipper interface {
	Zip(ctx context.Context)
	Unzip(ctx context.Context)
}
