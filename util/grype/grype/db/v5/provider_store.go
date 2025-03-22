package v5

import (
	"DIDTrustCore/util/grype/grype/match"
	"DIDTrustCore/util/grype/grype/vulnerability"
)

type ProviderStore struct {
	vulnerability.Provider
	match.ExclusionProvider
}
