package python

import (
	"DIDTrustCore/util/grype/grype/match"
	"DIDTrustCore/util/grype/grype/matcher/internal"
	"DIDTrustCore/util/grype/grype/pkg"
	"DIDTrustCore/util/grype/grype/vulnerability"
	syftPkg "github.com/anchore/syft/syft/pkg"
)

type Matcher struct {
	cfg MatcherConfig
}

type MatcherConfig struct {
	UseCPEs bool
}

func NewPythonMatcher(cfg MatcherConfig) *Matcher {
	return &Matcher{
		cfg: cfg,
	}
}

func (m *Matcher) PackageTypes() []syftPkg.Type {
	return []syftPkg.Type{syftPkg.PythonPkg}
}

func (m *Matcher) Type() match.MatcherType {
	return match.PythonMatcher
}

func (m *Matcher) Match(store vulnerability.Provider, p pkg.Package) ([]match.Match, []match.IgnoredMatch, error) {
	return internal.MatchPackageByEcosystemAndCPEs(store, p, m.Type(), m.cfg.UseCPEs)
}
