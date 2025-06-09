package main

import (
	"bytes"
	"image/jpeg"
	"log"
	"time"

	"github.com/kbinani/screenshot"
)

func startCaptureService() {
	ticker := time.NewTicker(time.Second / time.Duration(frameRate))
	defer ticker.Stop()

	for range ticker.C {
		img, err := screenshot.CaptureRect(screenBounds)
		if err != nil {
			log.Printf("截屏失败: %v", err)
			continue
		}

		buf := new(bytes.Buffer)
		if err := jpeg.Encode(buf, img, &jpeg.Options{Quality: quality}); err != nil {
			log.Printf("JPEG编码失败: %v", err)
			continue
		}

		updateFrameCache(buf.Bytes())
	}
}
