package main

import (
	"fmt"
	"net/http"
	"time"
)

func streamHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "multipart/x-mixed-replace; boundary=frame")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// 发送初始帧分隔符
	if _, err := w.Write([]byte("--frame\r\n")); err != nil {
		return
	}

	lastSent := time.Now()
	for {
		frame, frameTime := getFrameCache()

		// 只发送新帧
		if frameTime.After(lastSent) {
			if _, err := fmt.Fprintf(w,
				"Content-Type: image/jpeg\r\nContent-Length: %d\r\n\r\n",
				len(frame)); err != nil {
				return
			}

			if _, err := w.Write(frame); err != nil {
				return
			}

			if _, err := w.Write([]byte("\r\n--frame\r\n")); err != nil {
				return
			}

			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}

			lastSent = frameTime
		}

		// 防止CPU空转
		time.Sleep(time.Second / time.Duration(frameRate*2))

		select {
		case <-r.Context().Done():
			return
		default:
		}
	}
}
