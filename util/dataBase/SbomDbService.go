package dataBase

import (
	"DIDTrustCore/model"
	"time"
)

var Sbom_svc = NewSBOMManager(*Sbom_repo)

type sbomService struct {
	repo SbomRepo
}

func NewSBOMManager(repo SbomRepo) *sbomService {
	return &sbomService{repo: repo}
}

func (s *sbomService) GenerateSBOMRecord(userid uint, filename, url, format string) (*model.SBOMReport, error) {
	report := &model.SBOMReport{
		UserID:       userid,
		SBOMFilename: filename,
		DownloadURL:  url,
		Format:       format,
		ExpiresAt:    time.Now().Unix() + 24*3600,
	}

	if err := s.repo.Create(report); err != nil {
		return nil, err
	}
	return report, nil
}

func (s *sbomService) GetSBOMByID(id uint) (*model.SBOMReport, error) {
	return s.repo.GetByID(id)
}

func (s *sbomService) ListSBOMs(userID uint, page, size int) ([]model.SBOMReport, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 20
	}
	return s.repo.GetByUser(userID, page, size)
}

func (s *sbomService) RemoveSBOM(id uint) error {
	return s.repo.Delete(id)
}
