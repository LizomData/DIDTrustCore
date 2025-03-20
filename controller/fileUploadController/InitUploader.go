package fileUploadController

import (
	"fmt"
	"os"
)

var Uploader = initUploader()

// 上传结果结构体
type UploadResult struct {
	PublicURL string // 公开访问地址
	FileName  string // 存储的文件名
	FileSize  int64  // 文件大小(字节)
}

// 配置参数结构体
type UploadConfig struct {
	UploadDir    string   // 文件存储目录 (默认: ./uploads)
	PublicPath   string   // 公开访问路径 (默认: /uploads/)
	MaxFileSize  int64    // 最大文件大小 (默认: 32MB)
	AllowedTypes []string // 允许的文件类型 (默认: 全部)
}

type FileUploader struct {
	Config UploadConfig
}

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
