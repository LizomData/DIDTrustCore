package qualifier

import (
	"fmt"

	"DIDTrustCore/util/grype/grype/pkg/qualifier"
)

type Qualifier interface {
	fmt.Stringer
	Parse() qualifier.Qualifier
}
