package main

import (
	"sync"
	"time"
)

var (
	frameCache     []byte
	frameCacheTime time.Time
	cacheMutex     sync.RWMutex
)

// 更新帧缓存
func updateFrameCache(frame []byte) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	frameCache = frame
	frameCacheTime = time.Now()
}

// 获取帧缓存
func getFrameCache() ([]byte, time.Time) {
	cacheMutex.RLock()
	defer cacheMutex.RUnlock()
	return frameCache, frameCacheTime
}
