package model

type SBOMReport struct {
	ID           uint   `gorm:"primaryKey;comment:主键UUID" json:"sbomReportId"`
	UserID       uint   `gorm:"not null;index:idx_user;comment:关联用户ID" json:"-"`
	SBOMFilename string `gorm:"size:255;uniqueIndex:unq_sbom_file;comment:SBOM文件名" json:"sbom_filename"`
	DownloadURL  string `gorm:"type:varchar(1024);not null;comment:下载地址" json:"download_url"`
	Format       string `gorm:"type:varchar(20);check:format IN ('spdx-json','cyclonedx-json','syft-json');default:'spdx-json';comment:文件格式" json:"format"`
	// 使用 BIGINT 存储时间戳（秒级）
	ExpiresAt int64 `gorm:"type:bigint;index:idx_expires;comment:过期时间" json:"-"`
	CreatedAt int64 `gorm:"type:bigint;index:idx_created;comment:创建时间" json:"created_at"`
	UpdatedAt int64 `gorm:"type:bigint;index:idx_updated;comment:更新时间" json:"updated_at"`
}
