package fileUploadController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func upload(c *gin.Context) {

	// 获取上传文件
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请选择上传文件",
		})
		return
	}

	// 调用上传模块
	result, err := Uploader.UploadFile(fileHeader)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"url":  result.PublicURL,
			"name": result.FileName,
			"size": result.FileSize,
		},
	})

}

// 核心上传方法
func (u *FileUploader) UploadFile(fileHeader *multipart.FileHeader) (*UploadResult, error) {
	// 验证文件大小
	if fileHeader.Size > u.Config.MaxFileSize {
		return nil, fmt.Errorf("文件超过大小限制(最大 %dMB)", u.Config.MaxFileSize/(1<<20))
	}

	// 验证文件类型
	if len(u.Config.AllowedTypes) > 0 {
		ext := filepath.Ext(fileHeader.Filename)
		if !u.contains(u.Config.AllowedTypes, ext) {
			return nil, fmt.Errorf("不支持的文件类型: %s", ext)
		}
	}

	// 生成唯一文件名
	newFilename := fmt.Sprintf("%d_%d%s",
		time.Now().UnixNano(),
		time.Now().Nanosecond(),
		filepath.Ext(fileHeader.Filename))

	// 保存文件
	dstPath := filepath.Join(u.Config.UploadDir, newFilename)
	if err := saveUploadedFile(fileHeader, dstPath); err != nil {
		return nil, fmt.Errorf("文件保存失败: %w", err)
	}

	// 生成访问地址
	publicURL := filepath.Join(u.Config.PublicPath, newFilename)

	return &UploadResult{
		PublicURL: publicURL,
		FileName:  newFilename,
		FileSize:  fileHeader.Size,
	}, nil
}

// 辅助方法：检查切片包含
func (u *FileUploader) contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// 修改保存方法实现（不依赖 Gin 框架）
var saveUploadedFile = func(file *multipart.FileHeader, dst string) error {
	// 1. 打开上传文件
	src, err := file.Open()
	if err != nil {
		return fmt.Errorf("无法打开上传文件: %w", err)
	}
	defer src.Close()

	// 2. 创建目标文件
	out, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("无法创建目标文件: %w", err)
	}
	defer out.Close()

	// 3. 复制文件内容
	if _, err := io.Copy(out, src); err != nil {
		return fmt.Errorf("文件写入失败: %w", err)
	}

	return nil
}
