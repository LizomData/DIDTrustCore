package sbomController

import (
	"DIDTrustCore/controller/fileUploadController"
	"DIDTrustCore/util/extractorCustom"
	"DIDTrustCore/util/sbom"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type SBOMController struct {
	uploadDir  string
	publicPath string
}

// GenerateSBOMRequest 定义请求体结构
type GenerateSBOMRequest struct {
	FileURL string `json:"file_url"`
	Format  string `json:"format" enums:"spdx-json,cyclonedx-json,syft-json"`
}

// GenerateSBOM 处理SBOM生成请求
func generate(c *gin.Context) {
	var req GenerateSBOMRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    40010,
			"message": "无效的请求参数: " + err.Error(),
		})
		return
	}

	// 验证并解析文件URL
	filename, err := validateFileURL(req.FileURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    40011,
			"message": "文件URL无效: " + err.Error(),
		})
		return
	}

	// 获取本地文件路径
	filePath := filepath.Join(fileUploadController.Uploader.Config.UploadDir, filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    40401,
			"message": "文件不存在",
		})
		return
	}
	fmt.Println(filePath)

	// 创建解压器实例
	extractor := extractorCustom.Extractor{}

	// 解压文件到临时目录
	extractDir, err := extractor.Extract(filePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    40012,
			"message": "文件解压失败: " + err.Error(),
		})
		return
	}
	defer os.RemoveAll(extractDir) // 确保清理临时目录

	// 生成SBOM（使用解压后的目录）
	sbomBytes, err := sbom.GenerateSBOM(extractDir, req.Format)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    50001,
			"message": "SBOM生成失败: " + err.Error(),
		})
		return
	}
	c.Data(http.StatusOK, "application/json", sbomBytes)

}

// 验证文件URL并提取文件名
func validateFileURL(url string) (string, error) {
	if !strings.HasPrefix(url, fileUploadController.Uploader.Config.PublicPath) {
		return "", fmt.Errorf("非法的文件路径")
	}
	filename := strings.TrimPrefix(url, fileUploadController.Uploader.Config.PublicPath)
	if filename == "" || strings.Contains(filename, "..") {
		return "", fmt.Errorf("无效的文件名")
	}
	return filename, nil
}
