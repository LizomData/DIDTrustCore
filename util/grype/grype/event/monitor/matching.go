package monitor

import (
	"github.com/wagoodman/go-progress"

	"DIDTrustCore/util/grype/grype/vulnerability"
)

type Matching struct {
	PackagesProcessed progress.Progressable
	MatchesDiscovered progress.Monitorable
	Fixed             progress.Monitorable
	Ignored           progress.Monitorable
	Dropped           progress.Monitorable
	BySeverity        map[vulnerability.Severity]progress.Monitorable
}
