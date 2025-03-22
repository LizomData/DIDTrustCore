package pkgUploadController

import (
	"fmt"
	"os"
)

var Uploader = initUploader()

func initUploader() *FileUploader {
	// 初始化上传模块
	_uploader, err := NewUploader(UploadConfig{
		UploadDir:    "./uploads",
		PublicPath:   "/uploads/",
		MaxFileSize:  50 << 20, // 50MB
		AllowedTypes: []string{".zip", ".tar.gz"},
	})
	if err != nil {
		panic(fmt.Sprintf("初始化上传模块失败: %v", err))
	}
	return _uploader
}

// 初始化上传器
func NewUploader(cfg UploadConfig) (*FileUploader, error) {
	// 设置默认值
	if cfg.UploadDir == "" {
		cfg.UploadDir = "./uploads"
	}
	if cfg.PublicPath == "" {
		cfg.PublicPath = "/uploads/"
	}
	if cfg.MaxFileSize == 0 {
		cfg.MaxFileSize = 32 << 20 // 32MB
	}

	// 创建存储目录
	if err := os.MkdirAll(cfg.UploadDir, 0755); err != nil {
		return nil, fmt.Errorf("创建上传目录失败: %w", err)
	}

	return &FileUploader{Config: cfg}, nil
}
