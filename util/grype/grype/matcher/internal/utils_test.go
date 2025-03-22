package internal

import (
	"testing"

	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"

	"DIDTrustCore/util/grype/grype/match"
	"DIDTrustCore/util/grype/grype/vulnerability"
)

func assertMatchesUsingIDsForVulnerabilities(t testing.TB, expected, actual []match.Match) {
	t.Helper()
	require.Len(t, actual, len(expected))
	for idx, a := range actual {
		// only compare the vulnerability ID, nothing else
		a.Vulnerability = vulnerability.Vulnerability{Reference: vulnerability.Reference{ID: a.Vulnerability.ID}}
		for _, d := range deep.Equal(expected[idx], a) {
			t.Errorf("diff idx=%d: %+v", idx, d)
		}
	}
}
