/*
Package elixir provides a concrete Cataloger implementation relating to packages within the Elixir language ecosystem.
*/
package elixir

import (
	"DIDTrustCore/util/sbom/syft/syft/pkg"
	"DIDTrustCore/util/sbom/syft/syft/pkg/cataloger/generic"
)

// NewMixLockCataloger returns a cataloger object for Elixir mix.lock files.
func NewMixLockCataloger() pkg.Cataloger {
	return generic.NewCataloger("elixir-mix-lock-cataloger").
		WithParserByGlobs(parseMixLock, "**/mix.lock")
}
