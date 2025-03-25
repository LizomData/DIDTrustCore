package sbomDb

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

func (s *sbomService) GenerateSBOMRecord(userid uint, didid, filename, url, format string) (*model.SBOMReport, error) {
	report := &model.SBOMReport{
		UserID:       userid,
		DidID:        didid,
		SbomFilename: filename,
		DownloadURL:  url,
		Format:       format,
		ExpiresAt:    time.Now().Unix() + 24*3600,
	}

	if err := s.repo.Create(report); err != nil {
		return nil, err
	}
	return report, nil
}

func (s *sbomService) GetSBOMByDidID(didid string) (*model.SBOMReport, error) {
	return s.repo.GetByDidID(didid)
}
