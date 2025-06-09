package main

import (
	"github.com/kbinani/screenshot"
	"image"
	"log"
	"net/http"
	"os"
)

var (
	quality      = 50                     // JPEG质量(1-100)
	frameRate    = 10                     // 帧率(FPS)
	screenBounds = image.Rect(0, 0, 0, 0) // 屏幕尺寸
)

func main() {
	// 初始化截屏区域
	if n := screenshot.NumActiveDisplays(); n <= 0 {
		log.Fatal("未检测到活动显示器")
	}
	screenBounds = screenshot.GetDisplayBounds(0)

	// 启动截屏服务
	go startCaptureService()

	// 设置路由
	http.HandleFunc("/stream", streamHandler)

	wd, _ := os.Getwd()
	log.Printf("当前工作目录: %s GOPATH：%s", wd, os.Getenv("GOPATH"))
	// 使用 http.FileServer 提供整个静态资源目录
	fs := http.FileServer(http.Dir("./demo/static"))
	http.Handle("/", fs)

	/**
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("请求 / -> 提供 index.html")
		//name := "./demo/static/index.html" // 相对路径
		name := "E:/Git/go/screenshot/demo/static/index.html" // 绝对路径
		http.ServeFile(w, r, name)
	})
	*/

	// 启动服务器
	log.Println("服务启动: http://localhost:8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
