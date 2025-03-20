package extractorCustom

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Extractor 处理不同格式的压缩包解压
type Extractor struct{}

// Extract 根据文件类型解压到临时目录
func (e *Extractor) Extract(srcPath string) (string, error) {
	ext := filepath.Ext(srcPath)
	switch ext {
	case ".zip":
		return e.extractZip(srcPath)
	case ".tar.gz":
		return e.extractTarGz(srcPath)
	default:
		return "", fmt.Errorf("不支持的文件类型: %s", ext)
	}
}

// 解压ZIP文件
func (e *Extractor) extractZip(srcPath string) (string, error) {
	tempDir, err := os.MkdirTemp("", "sbom-extract-*")
	if err != nil {
		return "", fmt.Errorf("无法创建临时目录: %w", err)
	}

	r, err := zip.OpenReader(srcPath)
	if err != nil {
		os.RemoveAll(tempDir)
		return "", fmt.Errorf("无法打开ZIP文件: %w", err)
	}
	defer r.Close()

	for _, f := range r.File {
		err := extractZipFile(f, tempDir)
		if err != nil {
			os.RemoveAll(tempDir)
			return "", err
		}
	}

	return tempDir, nil
}

func extractZipFile(f *zip.File, dest string) error {
	// 防止路径遍历攻击
	path := filepath.Join(dest, f.Name)
	if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
		return fmt.Errorf("非法文件路径: %s", f.Name)
	}

	if f.FileInfo().IsDir() {
		return os.MkdirAll(path, 0755)
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}
	defer outFile.Close()

	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	_, err = io.Copy(outFile, rc)
	return err
}

// TODO: 添加tar.gz解压实现
func (e *Extractor) extractTarGz(srcPath string) (string, error) {
	return "", fmt.Errorf("tar.gz解压暂未实现")
}
