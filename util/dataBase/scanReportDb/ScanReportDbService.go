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

func (s *ScanReportDbService) CreateRecord(userid uint, didid, filename, download_url string) (*model.ScanReport, error) {
	report := &model.ScanReport{
		UserID:             userid,
		DidID:              didid,
		ScanResultFilename: filename,
		DownloadURL:        download_url,
		ExpiresAt:          time.Now().Unix() + 24*3600,
	}

	if err := s.repo.Create(report); err != nil {
		return nil, err
	}
	return report, nil
}

func (s *ScanReportDbService) GetRecordByDidID(didid string) (*model.ScanReport, error) {
	return s.repo.GetByDidID(didid)
}
