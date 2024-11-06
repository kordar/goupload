package goupload

type Bucket struct {
	Name   string      `json:"name" xml:"name"`     // bucket名称
	Driver string      `json:"driver" xml:"driver"` // bucket驱动 cos、oss、local、ipfs等
	Params interface{} `json:"params" xml:"params"` // bucket参数 {"id":"xxx","pwd":"123456"}
}

type BucketObject struct {
	Id           string      `json:"id" xml:"id"`
	Path         string      `json:"path" xml:"path"`                   // 存储路径
	LastModified string      `json:"last_modified" xml:"last_modified"` // 存储时间
	Size         int64       `json:"size" xml:"size"`                   // 文件大小
	FileType     string      `json:"file_type" xml:"file_type"`         // 文件类型 dir,file
	FileExt      string      `json:"file_ext" xml:"file_ext"`           // 文件扩展名
	ParentId     string      `json:"parent_id" xml:"parent_id"`         // 父目录
	Params       interface{} `json:"params" xml:"params"`
}

type BucketTreeObject struct {
	Item     BucketObject       `json:"item" xml:"item"`
	Children []BucketTreeObject `json:"children" xml:"children"`
}
