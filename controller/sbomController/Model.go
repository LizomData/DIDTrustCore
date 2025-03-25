package sbomController

// 配置参数结构体
type SbomConfig struct {
	SBOMStorageDir string // 文件存储目录 (默认: ./sbomStorage)
	PublicPath     string // 公开访问路径 (默认: /sbom_list/)

}

type SbomGenerator struct {
	Config SbomConfig
}

// GenerateSBOMRequest 定义请求体结构
type GenerateSBOMRequest struct {
	DidID      string `json:"didid"`
	PkgFileUrl string `json:"pkg_file_url"`
	Format     string `json:"format" enums:"spdx-json,cyclonedx-json,syft-json"`
}
