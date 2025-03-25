package scanReportDb

import (
	"DIDTrustCore/model"
	"DIDTrustCore/util/dataBase"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

var scanReportRepo = NewScanReportRepository(dataBase.Db)

type ScanReportRepo struct {
	db *gorm.DB
}

func NewScanReportRepository(db *gorm.DB) *ScanReportRepo {
	return &ScanReportRepo{db: db}
}

// Create 创建分析记录
func (r *ScanReportRepo) Create(report *model.ScanReport) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 参数校验
		if report.UserID == 0 || report.DownloadURL == "" {
			return errors.New("必要字段缺失：UserID/DownloadURL")
		}

		// 存在性检查
		var existing model.ScanReport
		err := tx.Where("did_id = ?", report.DidID).First(&existing).Error

		switch {
		case err == nil: // 记录存在，执行更新
			return tx.Model(&existing).Updates(map[string]interface{}{
				"scan_result_filename": report.ScanResultFilename,
				"download_url":         report.DownloadURL,
				"expires_at":           report.ExpiresAt,
				"updated_at":           time.Now().Unix(),
			}).Error

		case errors.Is(err, gorm.ErrRecordNotFound): // 新记录
			report.CreatedAt = time.Now().Unix()
			report.UpdatedAt = time.Now().Unix()
			return tx.Create(report).Error

		default: // 其他错误
			return fmt.Errorf("数据库查询失败: %w", err)
		}
	})
}

// GetByID 通过ID获取记录
func (r *ScanReportRepo) GetByDidID(didid string) (*model.ScanReport, error) {
	var report model.ScanReport
	err := r.db.Where("did_id = ?", didid).First(&report).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &report, nil
	}
	return &report, err
}
