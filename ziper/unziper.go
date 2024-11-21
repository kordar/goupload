package ziper

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func unzipToCallback(files []*zip.File, callback func(*zip.File)) error {
	for _, file := range files {
		callback(file)
	}
	return nil
}

func unzipToFile(files []*zip.File, dest string) error {
	return unzipToCallback(files, func(file *zip.File) {
		filename := filepath.Join(dest, file.Name)
		if file.FileInfo().IsDir() {
			_ = os.MkdirAll(filename, os.ModePerm)
			return
		}

		if err := os.MkdirAll(filepath.Dir(filename), os.ModePerm); err != nil {
			return
		}

		srcfile, err := file.Open()
		defer srcfile.Close()
		if err != nil {
			return
		}

		// 创建目标文件
		destfile, err := os.Create(filename)
		defer destfile.Close()
		if err != nil {
			return
		}

		_, _ = io.Copy(destfile, srcfile)

	})
}

func Unzip(src string, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return fmt.Errorf("unable to open ZIP file: %v", err)
	}
	defer r.Close()
	return unzipToFile(r.File, dest)
}

func UnzipWithBytes(b []byte, dest string) error {
	reader := bytes.NewReader(b)
	zipReader, err := zip.NewReader(reader, reader.Size())
	if err != nil {
		return fmt.Errorf("error creating zip reader: %v", err)
	}
	return unzipToFile(zipReader.File, dest)
}

func UnzipCallback(src string, callback func(*zip.File)) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return fmt.Errorf("unable to open ZIP file: %v", err)
	}
	defer r.Close()
	return unzipToCallback(r.File, callback)
}

func UnzipCallbackWithBytes(b []byte, callback func(*zip.File)) error {
	reader := bytes.NewReader(b)
	zipReader, err := zip.NewReader(reader, reader.Size())
	if err != nil {
		return fmt.Errorf("error creating zip reader: %v", err)
	}
	return unzipToCallback(zipReader.File, callback)
}
