package ziper

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func unzipToFile(files []*zip.File, dest string) error {
	for _, file := range files {
		filename := filepath.Join(dest, file.Name)
		if file.FileInfo().IsDir() {
			_ = os.MkdirAll(filename, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filename), os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}

		srcfile, err := file.Open()
		if err != nil {
			return fmt.Errorf("unable to open file: %v", err)
		}

		// 创建目标文件
		destfile, err := os.Create(filename)
		if err != nil {
			_ = srcfile.Close()
			return fmt.Errorf("unable to create file: %v", err)
		}

		// 将文件内容拷贝到目标文件
		if _, err := io.Copy(destfile, srcfile); err != nil {
			_ = srcfile.Close()
			_ = destfile.Close()
			return fmt.Errorf("file decompression failed: %v", err)
		}

		_ = srcfile.Close()
		_ = destfile.Close()
	}

	return nil
}

func Unzip(src string, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return fmt.Errorf("unable to open ZIP file: %v", err)
	}
	defer r.Close()
	return unzipToFile(r.File, dest)
}
