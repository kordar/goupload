package ziper

import (
	"archive/zip"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func AddReaderToZip(zipWriter *zip.Writer, reader io.Reader, filename string) error {
	// 创建文件头
	header := &zip.FileHeader{}
	header.Name = filename
	header.Method = zip.Deflate

	// 创建文件写入器
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return fmt.Errorf("unable to create file writer: %v", err)
	}

	// 将文件内容写入 ZIP
	_, err = io.Copy(writer, reader)
	if err != nil {
		return fmt.Errorf("failed to write file to ZIP: %v", err)
	}

	return nil
}

func AddFileToZip(zipWriter *zip.Writer, filename string) error {
	// 打开要添加的文件
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("can not open the file: %v", err)
	}
	defer file.Close()

	// 获取文件信息
	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("unable to obtain file information: %v", err)
	}

	// 创建文件头
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return fmt.Errorf("unable to create file header: %v", err)
	}
	header.Name = filepath.Base(filename)
	header.Method = zip.Deflate // 使用压缩算法

	// 创建文件写入器
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return fmt.Errorf("unable to create file writer: %v", err)
	}

	// 将文件内容写入 ZIP
	_, err = io.Copy(writer, file)
	if err != nil {
		return fmt.Errorf("failed to write file to ZIP: %v", err)
	}

	return nil
}

func CreateZip(zipFilename string, files []string) error {
	// 创建 ZIP 文件
	zipFile, err := os.Create(zipFilename)
	if err != nil {
		return fmt.Errorf("unable to create ZIP file: %v", err)
	}
	defer zipFile.Close()

	// 创建 ZIP 写入器
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 添加文件到 ZIP
	for _, file := range files {
		log.Printf("adding files: %s\n", file)
		if e := AddFileToZip(zipWriter, file); e != nil {
			return e
		}
	}

	return nil
}

func CreateZipWithReader(zipFilename string, readers map[string]io.Reader) error {
	// 创建 ZIP 文件
	zipFile, err := os.Create(zipFilename)
	if err != nil {
		return fmt.Errorf("unable to create ZIP file: %v", err)
	}
	defer zipFile.Close()

	// 创建 ZIP 写入器
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 添加文件到 ZIP
	for filename, reader := range readers {
		log.Printf("adding files: %s\n", filename)
		if e := AddReaderToZip(zipWriter, reader, filename); e != nil {
			return e
		}
	}

	return nil
}

func OutputCreateZip(files []string) ([]byte, error) {
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)
	zipWriter := zip.NewWriter(writer)
	defer zipWriter.Close()

	for _, file := range files {
		log.Printf("adding files: %s\n", file)
		if err := AddFileToZip(zipWriter, file); err != nil {
			return nil, err
		}
	}

	return buffer.Bytes(), nil
}

func OutputCreateZipWithReader(readers map[string]io.Reader) ([]byte, error) {
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)
	zipWriter := zip.NewWriter(writer)
	defer zipWriter.Close()

	// 添加文件到 ZIP
	for filename, reader := range readers {
		log.Printf("adding reader: %s\n", filename)
		if err := AddReaderToZip(zipWriter, reader, filename); err != nil {
			return nil, err
		}
	}

	return buffer.Bytes(), nil
}
