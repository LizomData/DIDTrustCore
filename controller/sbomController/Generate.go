package sbomController

import (
	"DIDTrustCore/common"
	"DIDTrustCore/controller/fileUploadController"
	"DIDTrustCore/model/requestBase"
	"DIDTrustCore/util/dataBase"
	"DIDTrustCore/util/extractorCustom"
	"DIDTrustCore/util/sbom"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// @Summary 生成SBOM接口
// @Description 根据已上传的软件包生成SBOM文件
// @Tags SBOM管理
// @Accept json
// @Produce json
// @Param Authorization	header		string	true	"jwt"
// @Param body body GenerateSBOMRequest true "生成参数"
// @Success 200 {object} GenerateSbomResult "SBOM清单信息"
// @Router /api/v1/sbom/generate [post]
func generate(c *gin.Context) {
	var req GenerateSBOMRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(requestBase.ResponseBody(requestBase.ParameterError, "无效的请求参数,检查软件包路径", gin.H{}))
		return
	}

	// 验证并解析文件URL
	filename, err := validateFileURL(req.FileURL)
	if err != nil {
		c.JSON(requestBase.ResponseBody(requestBase.FileNotFound, "文件URL无效,检查软件包路径", gin.H{}))
		return
	}

	// 获取本地文件路径
	filePath := filepath.Join(fileUploadController.Uploader.Config.UploadDir, filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(requestBase.ResponseBody(requestBase.FileNotFound, "软件包不存在,检查软件包上传状态", gin.H{}))
		return
	}

	// 创建解压器实例
	extractor := extractorCustom.Extractor{}

	// 解压文件到临时目录
	extractDir, err := extractor.Extract(filePath)
	if err != nil {
		c.JSON(requestBase.ResponseBody(requestBase.FileUnzipFailed, "软件包解压失败,检查软件包上传状态", gin.H{}))
		return
	}
	defer os.RemoveAll(extractDir) // 确保清理临时目录

	// 生成SBOM（使用解压后的目录）
	sbomBytes, err := sbom.GenerateSBOM(extractDir, req.Format)
	if err != nil {
		c.JSON(requestBase.ResponseBody(requestBase.SBOMFailed, "SBOM生成失败", gin.H{}))
		return
	}

	// 生成唯一文件名
	sbomFilename := fmt.Sprintf("sbom_%d_%d.%s.json", time.Now().UnixNano(),
		time.Now().Nanosecond(), strings.Split(req.Format, "-")[0])
	download_url := fmt.Sprintf("%s%s", Generator.Config.PublicPath, sbomFilename)

	// 保存到持久化存储
	sbomPath := filepath.Join(Generator.Config.SBOMStorageDir, sbomFilename)
	if err := os.WriteFile(sbomPath, sbomBytes, 0644); err != nil {
		c.JSON(requestBase.ResponseBody(
			requestBase.SBOMFailed,
			"SBOM生成失败: "+err.Error(),
			gin.H{}))
		return
	}
	user, err := common.GetUserFromContext(c)
	//保存到数据库,如果已经登陆了
	if _, err := dataBase.Sbom_svc.GenerateSBOMRecord(user.ID, sbomFilename, download_url, req.Format); err != nil {
		c.JSON(requestBase.ResponseBody(
			requestBase.SBOMFailed,
			"SBOM保存云端失败: "+err.Error(),
			gin.H{}))
		return
	}

	// 返回下载链接
	c.JSON(requestBase.ResponseBodySuccess(gin.H{
		"download_url": download_url,
	}))

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
