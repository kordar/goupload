package ziper

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"testing"
)

func TestCreateZip(t *testing.T) {
	CreateZip(
		"/Users/mac/Pictures/bucket/test/AA/33.zip",
		[]string{
			"/Users/mac/Pictures/bucket/test/AA/2.txt",
			"/Users/mac/Pictures/bucket/test/AA/4.txt",
		},
	)
}

func TestAddFileToZip(t *testing.T) {
	// 创建 ZIP 写入器
	zipFile, err := os.Open("/Users/mac/Pictures/bucket/test/AA/33.zip")
	log.Println("=====================", err)
	defer zipFile.Close()
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()
	AddFileToZip(zipWriter, "/Users/mac/Pictures/bucket/test/AA/5.txt")
	AddFileToZip(zipWriter, "/Users/mac/Pictures/bucket/test/AA/6.txt")
}

func TestCreateZipWithReader(t *testing.T) {
	zipFile, _ := os.Create("/Users/mac/Pictures/bucket/test/AA/44.zip")
	defer zipFile.Close()
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	open, _ := os.Open("/Users/mac/Pictures/bucket/test/AA/5.txt")
	defer open.Close()
	AddReaderToZip(zipWriter, open, "aa.txt")

	open2, _ := os.Open("/Users/mac/Pictures/bucket/test/AA/6.txt")
	defer open2.Close()
	AddReaderToZip(zipWriter, open2, "bb.txt")
}

func TestCreateZipWithReader2(t *testing.T) {
	open, _ := os.Open("/Users/mac/Pictures/bucket/test/AA/5.txt")
	defer open.Close()

	open2, _ := os.Open("/Users/mac/Pictures/bucket/test/AA/6.txt")
	defer open2.Close()
	// -----------------
	_ = CreateZipWithReader("/Users/mac/Pictures/bucket/test/AA/33.zip", map[string]io.Reader{
		"1.txt": open, "2.txt": open2,
	})
}

func TestUnzip(t *testing.T) {
	Unzip("/Users/mac/Pictures/bucket/test/AA/33.zip", "/Users/mac/Pictures/bucket/test/AA/target")
}
