package pkgUploadController

// 上传结果结构体
type UploadResult struct {
	PublicURL string // 公开访问地址
	FileName  string // 存储的文件名
	FileSize  int64  // 文件大小(字节)
}

// 配置参数结构体
type UploadConfig struct {
	UploadDir    string   // 文件存储目录 (默认: ./uploads)
	PublicPath   string   // 公开访问路径 (默认: /uploads/)
	MaxFileSize  int64    // 最大文件大小 (默认: 32MB)
	AllowedTypes []string // 允许的文件类型 (默认: 全部)
}

type FileUploader struct {
	Config UploadConfig
}
type QueryRequest struct {
	Page int `json:"page" `
	Size int `json:"size"`
}

type GetDetailRequest struct {
	DidID string `json:"didid"`
}
type GetDetailResponse struct {
	SbomFileUrl       string `json:"sbomFileUrl"`
	ScanReportFileUrl string `json:"scanReportFileUrl"`
}
