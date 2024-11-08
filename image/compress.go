package image

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/jpeg"
	"io"
	"os"
)

func saveCompressOutputFile(img image.Image, outputpath string, quality int, width, height int, filter imaging.ResampleFilter) error {
	outfile, err := os.Create(outputpath)
	if err != nil {
		return fmt.Errorf("failed to mkdir: %v", err)
	}
	defer outfile.Close()
	// resizedImg := imaging.Resize(img, width, height, imaging.Lanczos)
	if width == 0 && height == 0 {
		opts := &jpeg.Options{Quality: quality}
		return jpeg.Encode(outfile, img, opts)
	}

	resizedimg := imaging.Resize(img, width, height, filter)
	opts := &jpeg.Options{Quality: quality}
	return jpeg.Encode(outfile, resizedimg, opts)
}

func outputCompressData(img image.Image, quality int, width, height int, filter imaging.ResampleFilter) ([]byte, error) {
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)

	if width == 0 && height == 0 {
		opts := &jpeg.Options{Quality: quality}
		if err := jpeg.Encode(writer, img, opts); err != nil {
			return nil, err
		} else {
			_ = writer.Flush()
			return buffer.Bytes(), nil
		}
	}

	resizedimg := imaging.Resize(img, width, height, filter)

	opts := &jpeg.Options{Quality: quality}
	if err := jpeg.Encode(writer, resizedimg, opts); err != nil {
		return nil, fmt.Errorf("failed to encoding image: %v", err)
	}
	_ = writer.Flush()
	return buffer.Bytes(), nil
}

func SaveCompressAndResize(inputPath, outputPath string, quality int, width, height int, filter imaging.ResampleFilter) error {
	img, err := imaging.Open(inputPath)
	if err != nil {
		return fmt.Errorf("can't open the image: %v", err)
	}
	return saveCompressOutputFile(img, outputPath, quality, width, height, filter)
}

func SaveCompressAndResizeByReader(inputReader io.Reader, outputPath string, quality int, width, height int, filter imaging.ResampleFilter) error {
	img, err := imaging.Decode(inputReader)
	if err != nil {
		return fmt.Errorf("can't open the image: %v", err)
	}
	return saveCompressOutputFile(img, outputPath, quality, width, height, filter)
}

func OutputCompressAndResize(inputPath string, quality int, width, height int, filter imaging.ResampleFilter) ([]byte, error) {
	img, err := imaging.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("can't open the image: %v", err)
	}
	return outputCompressData(img, quality, width, height, filter)
}

func OutputCompressAndResizeByReader(inputReader io.Reader, quality int, width, height int, filter imaging.ResampleFilter) ([]byte, error) {
	img, err := imaging.Decode(inputReader)
	if err != nil {
		return nil, fmt.Errorf("can't open the image: %v", err)
	}
	return outputCompressData(img, quality, width, height, filter)
}
