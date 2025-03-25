package pkgUploadController

import (
	"DIDTrustCore/common"
	"DIDTrustCore/model/requestBase"
	"DIDTrustCore/util"
	pkgDB "DIDTrustCore/util/dataBase/pkgDb"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// 文件上传接口
// @Summary 上传软件包
// @Description 上传软件包压缩包到服务器并返回访问地址,支持.zip和.tar.gz格式压缩包,格式采用multipart/form-data,字段为file
// @Tags 软件包管理
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "选择要上传的文件"
// @Success 200 {object} model.PkgRecord "上传成功"
// @Router /api/v1/pkg/upload [post]
func upload(c *gin.Context) {

	// 获取上传文件
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(requestBase.ResponseBody(requestBase.FileNotFound, "请选择上传文件", gin.H{}))
		return
	}

	// 调用上传模块
	result, err := Uploader.UploadFile(fileHeader)
	if err != nil {
		c.JSON(requestBase.ResponseBody(requestBase.UploadFailed, "上传文件失败:"+err.Error(), gin.H{}))
		return
	}

	user, err := common.GetUserFromContext(c)
	if err != nil {
		c.JSON(requestBase.ResponseBody(requestBase.UploadFailed, "获取用户id失败:"+err.Error(), gin.H{}))
		return
	}
	//保存上传记录
	record, err := pkgDB.Svc.CreateRecord(user.ID, result.FileName, result.PublicURL)
	if err != nil {
		c.JSON(requestBase.ResponseBody(requestBase.UploadFailed, "保存上传文件记录失败:"+err.Error(), gin.H{}))
		return
	}

	// 返回成功响应
	c.JSON(requestBase.ResponseBodySuccess(record))

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
	newFilename := util.GenerateUniqueFilename(fileHeader.Filename)

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
