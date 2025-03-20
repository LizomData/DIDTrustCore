package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
)

func SaveUploadedFile(c *gin.Context, file *multipart.FileHeader) (string, error) {
	// 创建临时目录
	tempDir, err := os.MkdirTemp("", "sbom-upload-*")
	if err != nil {
		return "", err
	}

	// 保存文件
	dstPath := filepath.Join(tempDir, file.Filename)
	if err := c.SaveUploadedFile(file, dstPath); err != nil {
		return "", err
	}

	return tempDir, nil
}

func CloneGitHubRepo(repoURL string) (string, error) {
	// 创建临时目录
	tempDir, err := os.MkdirTemp("", "sbom-repo-*")
	if err != nil {
		return "", err
	}

	// 执行 git clone
	cmd := exec.Command("git", "clone", repoURL, tempDir)
	if output, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("git clone failed: %s", string(output))
	}

	return tempDir, nil
}

func CleanupTempDir(dir string) {
	os.RemoveAll(dir)
}

func SaveUploadedFileWithCleanup(c *gin.Context, file *multipart.FileHeader) (string, func(), error) {
	tempDir, err := os.MkdirTemp("", "sbom-upload-*")
	if err != nil {
		return "", nil, fmt.Errorf("创建临时目录失败: %w", err)
	}

	cleanup := func() { os.RemoveAll(tempDir) }

	if err := c.SaveUploadedFile(file, filepath.Join(tempDir, file.Filename)); err != nil {
		cleanup()
		return "", nil, fmt.Errorf("保存文件失败: %w", err)
	}

	return tempDir, cleanup, nil
}

// util/git.go
func CloneGitRepoWithCleanup(repoURL string) (string, func(), error) {
	tempDir, err := os.MkdirTemp("", "sbom-repo-*")
	if err != nil {
		return "", nil, fmt.Errorf("创建临时目录失败: %w", err)
	}

	cleanup := func() { os.RemoveAll(tempDir) }

	cmd := exec.Command("git", "clone", repoURL, tempDir)
	if output, err := cmd.CombinedOutput(); err != nil {
		cleanup()
		return "", nil, fmt.Errorf("克隆失败: %s\n%s", err, string(output))
	}

	return tempDir, cleanup, nil
}
