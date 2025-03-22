package stock

import (
	"strings"

	grypePkg "DIDTrustCore/util/grype/grype/pkg"
)

type Resolver struct {
}

func (r *Resolver) Normalize(name string) string {
	return strings.ToLower(name)
}

func (r *Resolver) Resolve(p grypePkg.Package) []string {
	return []string{r.Normalize(p.Name)}
}
