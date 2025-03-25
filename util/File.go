package util

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"strings"
	"time"
)

func GenerateUniqueFilename(originalName string) string {
	// 获取文件扩展名（包含点，如 ".jpg"）
	ext := filepath.Ext(originalName)

	// 去除扩展名的文件名主体（如 "photo"）
	nameWithoutExt := strings.TrimSuffix(originalName, ext)

	// 生成唯一标识（时间戳 + 随机数防碰撞）
	uniquePart := fmt.Sprintf("%d_%d", time.Now().UnixNano(), rand.Intn(1000))

	// 重组文件名：主体_唯一标识.扩展名
	return fmt.Sprintf("%s_%s%s", nameWithoutExt, uniquePart, ext)
}
