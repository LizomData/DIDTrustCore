package portage

import (
	"DIDTrustCore/util/grype/grype/match"
	"DIDTrustCore/util/grype/grype/matcher/internal"
	"DIDTrustCore/util/grype/grype/pkg"
	"DIDTrustCore/util/grype/grype/vulnerability"
	syftPkg "github.com/anchore/syft/syft/pkg"
)

type Matcher struct {
}

func (m *Matcher) PackageTypes() []syftPkg.Type {
	return []syftPkg.Type{syftPkg.PortagePkg}
}

func (m *Matcher) Type() match.MatcherType {
	return match.PortageMatcher
}

func (m *Matcher) Match(store vulnerability.Provider, p pkg.Package) ([]match.Match, []match.IgnoredMatch, error) {
	return internal.MatchPackageByDistro(store, p, m.Type())
}
