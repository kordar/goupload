package goupload

import (
	"fmt"
	"log"
	"testing"
)

func TestUploadManager_Buckets(t *testing.T) {
	client := NewCOSClient("liangxiang-1257614471", "ap-shanghai",
		"AKIDmNhma7MK2TdW8TX1KPFqBH6evwlKcePt", "FD0YrtjJdjaBWbZHjvrY0C8sNcbLr9Bx")
	manager := NewUploadManager(client)
	fmt.Println(manager.Buckets())
	list := manager.List("data/avatar/default/", "", 3)
	fmt.Println(list)
	for _, object := range list.Dirs {
		log.Println(object)
	}
}
