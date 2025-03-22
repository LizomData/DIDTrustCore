package scanReportDb

import (
	"DIDTrustCore/model"
	"DIDTrustCore/util/dataBase"
	"errors"
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
		// 验证必要字段
		if report.UserID == 0 || report.DownloadURL == "" {
			return errors.New("missing required fields")
		}
		report.CreatedAt = time.Now().Unix()
		report.UpdatedAt = time.Now().Unix()
		return tx.Create(report).Error
	})
}

// GetByID 通过ID获取记录
func (r *ScanReportRepo) GetByID(id uint) (*model.ScanReport, error) {
	var report model.ScanReport
	err := r.db.Where("id = ?", id).First(&report).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("record not found")
	}
	return &report, err
}

// Delete 删除记录（硬删除）
func (r *ScanReportRepo) Delete(id uint) error {
	return r.db.Unscoped().Delete(&model.ScanReport{ID: id}).Error
}
