package pkgDb

import (
	"DIDTrustCore/model"
	"DIDTrustCore/util/dataBase"
	"errors"
	"fmt"
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
func (r *PkgRepo) GetByDidID(didid string) (*model.PkgRecord, error) {
	var report model.PkgRecord
	err := r.db.Where("did_id = ?", didid).First(&report).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("record not found")
	}
	return &report, err
}
func (r *PkgRepo) GetByFilename(filename string) (*model.PkgRecord, error) {
	var report model.PkgRecord
	err := r.db.Where("pkg_filename = ?", filename).First(&report).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("record not found")
	}
	return &report, err
}

// 分页获取用户记录
func (r *PkgRepo) GetByUser(userID uint, page, size int) ([]model.PkgRecord, error) {
	var reports []model.PkgRecord
	offset := (page - 1) * size
	result := r.db.Where("user_id = ?", userID).
		Offset(offset).
		Limit(size).
		Order("created_at DESC").
		Find(&reports)
	return reports, result.Error
}

// UpdateDidIDByFilename 通过文件名更新DID标识
func (r *PkgRepo) UpdateDidIDByFilename(filename, didID string) error {
	// 参数校验
	if filename == "" {
		return errors.New("文件名不能为空")
	}
	if didID == "" {
		return errors.New("DID标识不能为空")
	}

	// 构建更新数据（包含更新时间戳）
	updates := map[string]interface{}{
		"did_id":     didID,
		"updated_at": time.Now().Unix(),
	}

	// 执行更新操作（使用主键字段）
	result := r.db.Model(&model.PkgRecord{}).
		Where("pkg_filename = ?", filename). // 主键条件
		Updates(updates)

	// 错误处理
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("记录不存在或无需更新")
	}
	return nil
}

// ClearDidIDByDID 通过现有DID标识清空该字段
func (r *PkgRepo) ClearDidIDByDID(didID string) error {
	// 参数校验
	if didID == "" {
		return errors.New("DID标识不能为空")
	}

	// 构建更新数据（含空值校验）
	updates := map[string]interface{}{
		"did_id":     "", // 设置为空字符串
		"updated_at": time.Now().Unix(),
	}

	// 执行更新操作
	result := r.db.Model(&model.PkgRecord{}).
		Where("did_id = ?", didID).
		Updates(updates)

	// 错误处理
	switch {
	case result.Error != nil:
		return fmt.Errorf("数据库操作失败: %w", result.Error)
	case result.RowsAffected == 0:
		return errors.New("未找到匹配的记录")
	default:
		return nil
	}
}
