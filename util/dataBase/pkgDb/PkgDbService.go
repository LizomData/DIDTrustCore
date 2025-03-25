package pkgDb

import (
	"DIDTrustCore/model"
	"errors"
	"fmt"
	"time"
)

var Svc = newSvc(*pkgRepo)

type PkgDbService struct {
	repo PkgRepo
}

func newSvc(repo PkgRepo) *PkgDbService {
	return &PkgDbService{repo: repo}
}

func (s *PkgDbService) CreateRecord(userid uint, filename, download_url string) (*model.PkgRecord, error) {
	record := &model.PkgRecord{
		UserID:      userid,
		PkgFilename: filename,
		DownloadURL: download_url,
		ExpiresAt:   time.Now().Unix() + 24*3600,
	}

	if err := s.repo.Create(record); err != nil {
		return nil, err
	}
	return record, nil
}

func (s *PkgDbService) GetRecordByDidID(didid string) (*model.PkgRecord, error) {
	return s.repo.GetByDidID(didid)
}

func (s *PkgDbService) ListPkgs(userID uint, page, size int) ([]model.PkgRecord, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 20
	}
	return s.repo.GetByUser(userID, page, size)
}

func (s *PkgDbService) UpdateRecordDidID(filename, didid string) error {
	// 参数校验
	if filename == "" {
		return errors.New("文件名不能为空")
	}
	if didid == "" {
		return errors.New("DID标识不能为空")
	}

	// 执行更新
	if err := s.repo.UpdateDidIDByFilename(filename, didid); err != nil {
		return fmt.Errorf("更新DID失败: %w", err)
	}

	return nil
}
