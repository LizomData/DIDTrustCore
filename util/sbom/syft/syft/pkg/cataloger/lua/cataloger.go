/*
Package lua provides a concrete Cataloger implementation for packages relating to the Lua language ecosystem.
*/
package lua

import (
	"DIDTrustCore/util/sbom/syft/syft/pkg"
	"DIDTrustCore/util/sbom/syft/syft/pkg/cataloger/generic"
)

// NewPackageCataloger returns a new cataloger object for Lua ROck.
func NewPackageCataloger() pkg.Cataloger {
	return generic.NewCataloger("lua-rock-cataloger").
		WithParserByGlobs(parseRockspec, "**/*.rockspec")
}
