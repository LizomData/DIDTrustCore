package sbomDb

import (
	"DIDTrustCore/model"
	"DIDTrustCore/util/dataBase"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

var Sbom_repo = NewSBOMRepository(dataBase.Db)

type SbomRepo struct {
	db *gorm.DB
}

func NewSBOMRepository(db *gorm.DB) *SbomRepo {
	return &SbomRepo{db: db}
}

// Create 创建SBOM记录
func (r *SbomRepo) Create(report *model.SBOMReport) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 参数校验
		if report.UserID == 0 || report.DownloadURL == "" || report.SbomFilename == "" {
			return errors.New("缺失必要字段：UserID/DownloadURL/SbomFilename")
		}

		// 存在性检查
		var existing model.SBOMReport
		err := tx.Where("did_id = ?", report.DidID).First(&existing).Error

		switch {
		case err == nil: // 记录存在，执行更新
			return tx.Model(&existing).Updates(map[string]interface{}{
				"sbom_filename": report.SbomFilename,
				"download_url":  report.DownloadURL,
				"format":        report.Format,
				"expires_at":    report.ExpiresAt,
				"updated_at":    time.Now().Unix(),
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
func (r *SbomRepo) GetByDidID(didid string) (*model.SBOMReport, error) {
	var report model.SBOMReport
	err := r.db.Where("did_id = ?", didid).First(&report).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &report, nil
	}
	return &report, err
}
