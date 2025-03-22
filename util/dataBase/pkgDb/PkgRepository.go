package pkgDB

import (
	"DIDTrustCore/model"
	"DIDTrustCore/util/dataBase"
	"errors"
	"gorm.io/gorm"
	"time"
)

var pkgRepo = NewPkgRepository(dataBase.Db)

type PkgRepo struct {
	db *gorm.DB
}

func NewPkgRepository(db *gorm.DB) *PkgRepo {
	return &PkgRepo{db: db}
}

// Create 创建分析记录
func (r *PkgRepo) Create(record *model.PkgRecord) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 验证必要字段
		if record.UserID == 0 || record.DownloadURL == "" {
			return errors.New("missing required fields")
		}
		record.CreatedAt = time.Now().Unix()
		record.UpdatedAt = time.Now().Unix()
		return tx.Create(record).Error
	})
}

// GetByID 通过ID获取记录
func (r *PkgRepo) GetByID(id uint) (*model.PkgRecord, error) {
	var report model.PkgRecord
	err := r.db.Where("id = ?", id).First(&report).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("record not found")
	}
	return &report, err
}

// Delete 删除记录（硬删除）
func (r *PkgRepo) Delete(id uint) error {
	return r.db.Unscoped().Delete(&model.PkgRecord{ID: id}).Error
}
