package model

type ScanReport struct {
	DidID              string `gorm:"type:varchar(255);not null;" json:"didid"`
	ScanResultFilename string `gorm:"type:varchar(255);not null;;comment:扫描结果文件名" json:"-"`
	UserID             uint   `gorm:"not null;index:idx_user;comment:关联用户ID" json:"-"`
	DownloadURL        string `gorm:"type:varchar(1024);not null;comment:下载地址" json:"download_url"`
	// 使用 BIGINT 存储时间戳（秒级）
	ExpiresAt int64 `gorm:"type:bigint;index:idx_expires;comment:过期时间" json:"-"`
	CreatedAt int64 `gorm:"type:bigint;index:idx_created;comment:创建时间" json:"created_at"`
	UpdatedAt int64 `gorm:"type:bigint;index:idx_updated;comment:更新时间" json:"updated_at"`
}
