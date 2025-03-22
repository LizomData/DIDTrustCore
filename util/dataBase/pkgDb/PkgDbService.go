package pkgDB

import (
	"DIDTrustCore/model"
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

func (s *PkgDbService) GetRecordByid(id uint) (*model.PkgRecord, error) {
	return s.repo.GetByID(id)
}

func (s *PkgDbService) RemoveRecord(id uint) error {
	return s.repo.Delete(id)
}
