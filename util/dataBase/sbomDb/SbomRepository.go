package sbomDb

import (
	"DIDTrustCore/model"
	"DIDTrustCore/util/dataBase"
	"errors"
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
		// 验证必要字段
		if report.UserID == 0 || report.DownloadURL == "" || report.SbomFilename == "" {
			return errors.New("missing required fields")
		}
		report.CreatedAt = time.Now().Unix()
		report.UpdatedAt = time.Now().Unix()
		return tx.Create(report).Error
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
