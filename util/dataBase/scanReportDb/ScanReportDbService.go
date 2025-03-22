package scanReportDb

import (
	"DIDTrustCore/model"
	"time"
)

var Svc = newSvc(*scanReportRepo)

type ScanReportDbService struct {
	repo ScanReportRepo
}

func newSvc(repo ScanReportRepo) *ScanReportDbService {
	return &ScanReportDbService{repo: repo}
}

func (s *ScanReportDbService) CreateRecord(userid uint, filename, download_url string) (*model.ScanReport, error) {
	report := &model.ScanReport{
		UserID:             userid,
		ScanResultFilename: filename,
		DownloadURL:        download_url,
		ExpiresAt:          time.Now().Add(24 * time.Hour),
	}

	if err := s.repo.Create(report); err != nil {
		return nil, err
	}
	return report, nil
}

func (s *ScanReportDbService) GetReportyID(id uint) (*model.ScanReport, error) {
	return s.repo.GetByID(id)
}

func (s *ScanReportDbService) RemoveRecord(id uint) error {
	return s.repo.Delete(id)
}
