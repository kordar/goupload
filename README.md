# goupload

定义`BucketUploader`接口，统一`bucket`操作界面，实现文件上传、下载、目录列表展示等功能。

## 接口定义

```go
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
    List(ctx context.Context, dir string, next interface{}, limit int, subCount bool, args ...interface{}) ([]BucketObject, interface{})
    
    // Count retrieves a list of BucketObject from the specified directory and returns an additional result, typically used for metadata or supplementary information.
    Count(ctx context.Context, dir string, args ...interface{}) int
    
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
    Tree(ctx context.Context, dir string, next interface{}, limit int, dep int, maxDep int, noleaf bool, subCount bool, args ...interface{}) []BucketTreeObject
    
    // Append appends data from an io.Reader to an existing object starting at the specified position.
    Append(ctx context.Context, name string, position int, fd io.Reader, args ...interface{}) (int, error)
    
    // AppendString appends data from a string to an existing object starting at the specified position.
    AppendString(ctx context.Context, name string, position int, content string, args ...interface{}) (int, error)
}
```

## 通过`UploaderManager`管理`BucketUploader`

- 实现并注册上传句柄到管理容器

```go
// 创建上传管理容器
func NewManager() *UploaderManager
// 创建上传管理容器并添加上传句柄
func NewManagerWithUploader(handlers ...BucketUploader) *UploaderManager 
// 添加上传句柄到管理容器
func (mgr *UploaderManager) Add(uploader ...BucketUploader)
```

- 快速使用
```go
// 获取buckets列表
func (mgr *UploaderManager) Buckets(bucketName string, args ...interface{}) []Bucket 
// 通过文件路径上传文件到bucket
func (mgr *UploaderManager) PutFromFile(bucketName string, name string, filePath string, args ...interface{}) error 
// 比特流上传到bucket
func (mgr *UploaderManager) Put(bucketName string, name string, fd io.Reader, args ...interface{}) error 
// 上传字符串（base64编码文件）到bucket
func (mgr *UploaderManager) PutString(bucketName string, name string, content string, args ...interface{}) error 
// 展示目录下的文件列表，subCount控制是否仅展示目录结构
func (mgr *UploaderManager) List(bucketName string, dir string, next interface{}, limit int, subCount bool, args ...interface{}) ([]BucketObject, interface{}) 
// 目录下的文件数量
func (mgr *UploaderManager) Count(bucketName string, dir string, args ...interface{}) int
// 删除文件
func (mgr *UploaderManager) Del(bucketName string, name string, args ...interface{}) error 
// 删除目录
func (mgr *UploaderManager) DelAll(bucketName string, dir string) 
// 批量删除
func (mgr *UploaderManager) DelMulti(bucketName string, objects []BucketObject) error 
// 获取文件对象
func (mgr *UploaderManager) Get(bucketName string, name string, args ...interface{}) ([]byte, error) 
// 获取文件并保存到目标路径
func (mgr *UploaderManager) GetToFile(bucketName string, name string, localPath string, args ...interface{}) error 
// 文件是否存在
func (mgr *UploaderManager) IsExist(bucketName string, name string) (bool, error) 
// 拷贝文件
func (mgr *UploaderManager) Copy(bucketName string, dest string, source string, args ...interface{}) error 
// 移动文件
func (mgr *UploaderManager) Move(bucketName string, dest string, source string, args ...interface{}) error 
// 重名文件
func (mgr *UploaderManager) Rename(bucketName string, dest string, source string, args ...interface{}) error 
// 打印文件树形结构，可控制深度
func (mgr *UploaderManager) Tree(bucketName string, path string, next interface{}, limit int, dep int, maxDep int, noleaf bool, subCount bool) []BucketTreeObject 
// 追加内容到文件
func (mgr *UploaderManager) Append(bucketName string, name string, position int, r io.Reader, args ...interface{}) (int, error) 
// 通过字符串追加内容到文件
func (mgr *UploaderManager) AppendString(bucketName string, name string, position int, content string, args ...interface{}) (int, error) 
```

## 常用工具

### 图片压缩

- 保存压缩文件到目标地址

```go
// width&height=0 仅降低图片质量，否则通过宽高比进行压缩
func SaveCompressAndResize(inputPath, outputPath string, quality int, width, height int, filter imaging.ResampleFilter) error
// 输入io.Reader对象进行压缩
func SaveCompressAndResizeByReader(inputReader io.Reader, outputPath string, quality int, width, height int, filter imaging.ResampleFilter)
```

- 返回压缩后的字节数组

```go
func OutputCompressAndResize(inputPath string, quality int, width, height int, filter imaging.ResampleFilter) ([]byte, error)
func OutputCompressAndResizeByReader(inputReader io.Reader, quality int, width, height int, filter imaging.ResampleFilter) ([]byte, error) 
```

### `zip`解压缩

- 创建`zip`包并保存到本地路径

```go
func CreateZip(zipFilename string, files []string) error
func CreateZipWithReader(zipFilename string, readers map[string]io.Reader) error
```

- 创建`zip`包并返回字节流

```go
func OutputCreateZip(files []string) ([]byte, error)
func OutputCreateZipWithReader(readers map[string]io.Reader) ([]byte, error)
```

### 解压`zip`

```go
func Unzip(src string, dest string) error
func UnzipWithBytes(b []byte, dest string) error
func UnzipCallback(src string, callback func(*zip.File)) error
func UnzipCallbackWithBytes(b []byte, callback func(*zip.File)) error
```
