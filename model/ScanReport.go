package model

import "time"

type ScanReport struct {
	ID                 uint      `gorm:"primaryKey;comment:主键UUID" json:"scanReportId"`
	UserID             uint      `gorm:"not null;index:idx_user;comment:关联用户ID" json:"-"`
	ScanResultFilename string    `gorm:"size:255;uniqueIndex:unq_sbom_file;comment:扫描结果文件名" json:"scanResultFilename"`
	DownloadURL        string    `gorm:"type:varchar(1024);not null;comment:下载地址" json:"download_url"`
	ExpiresAt          time.Time `gorm:"type:datetime;index:idx_expires;default:CURRENT_TIMESTAMP;comment:过期时间" json:"expires_at"`
	CreatedAt          time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`
	UpdatedAt          time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`
}
