package image

import (
	"github.com/disintegration/imaging"
	"testing"
)

func TestSaveCompressAndResize(t *testing.T) {
	_ = SaveCompressAndResize(
		"/Users/mac/Pictures/bucket/WechatIMG1105.jpg",
		"/Users/mac/Pictures/bucket/111233.jpg",
		30,
		0,
		0,
		imaging.Lanczos,
	)
}
