package sbomController

import (
	"fmt"
	"os"
)

var Generator = initSbomGenerator()

func initSbomGenerator() *SbomGenerator {
	// 初始化上传模块
	_generator, err := newGenerator(SbomConfig{})
	if err != nil {
		panic(fmt.Sprintf("初始化上传模块失败: %v", err))
	}
	return _generator
}

// 初始化
func newGenerator(cfg SbomConfig) (*SbomGenerator, error) {
	// 设置默认值
	if cfg.SBOMStorageDir == "" {
		cfg.SBOMStorageDir = "./tmp/sbomStorage"
	}
	if cfg.PublicPath == "" {
		cfg.PublicPath = "/sbom_list/"
	}

	// 创建存储目录
	if err := os.MkdirAll(cfg.SBOMStorageDir, 0755); err != nil {
		return nil, fmt.Errorf("创建上传目录失败: %w", err)
	}

	return &SbomGenerator{Config: cfg}, nil
}
