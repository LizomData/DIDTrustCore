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
	FileURL string `json:"file_url"`
	Format  string `json:"format" enums:"spdx-json,cyclonedx-json,syft-json"`
}

type QuerySBOMRequest struct {
	Page int `json:"page"` //页面索引
	Size int `json:"size"` //页面大小
}
type GenerateSbomResult struct {
	DownloadUrl string `json:"download_url"` // 下载地址
}
type QuerySbomResult struct {
	SbomReports string `json:"sbomReports" example:{}` // sbom记录
}
