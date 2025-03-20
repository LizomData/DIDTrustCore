package dataBase

import (
	"DIDTrustCore/model"
	"errors"
	"gorm.io/gorm"
)

type SbomRepo struct {
	db *gorm.DB
}

func NewSBOMRepository(db *gorm.DB) *SbomRepo {
	return &SbomRepo{db: db}
}

// Create 创建SBOM记录
func (r *SbomRepo) Create(report *model.SBOMReport) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 验证必要字段
		if report.UserID == 0 || report.DownloadURL == "" {
			return errors.New("missing required fields")
		}
		return tx.Create(report).Error
	})
}

// GetByID 通过ID获取记录
func (r *SbomRepo) GetByID(id uint) (*model.SBOMReport, error) {
	var report model.SBOMReport
	err := r.db.Where("id = ?", id).First(&report).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("record not found")
	}
	return &report, err
}

// GetByUser 分页获取用户记录
func (r *SbomRepo) GetByUser(userID uint, page, size int) ([]model.SBOMReport, error) {
	var reports []model.SBOMReport
	offset := (page - 1) * size
	result := r.db.Where("user_id = ?", userID).
		Offset(offset).
		Limit(size).
		Order("created_at DESC").
		Find(&reports)
	return reports, result.Error
}

// Delete 删除记录（硬删除）
func (r *SbomRepo) Delete(id uint) error {
	return r.db.Unscoped().Delete(&model.SBOMReport{ID: id}).Error
}
