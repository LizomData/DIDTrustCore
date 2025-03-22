package qualifier

import (
	"DIDTrustCore/util/grype/grype/pkg"
)

type Qualifier interface {
	Satisfied(p pkg.Package) (bool, error)
}
