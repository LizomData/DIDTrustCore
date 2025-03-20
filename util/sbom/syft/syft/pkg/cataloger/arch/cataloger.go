/*
Package arch provides a concrete Cataloger implementations for packages relating to the Arch linux distribution.
*/
package arch

import (
	"DIDTrustCore/util/sbom/syft/syft/pkg"
	"DIDTrustCore/util/sbom/syft/syft/pkg/cataloger/generic"
	"DIDTrustCore/util/sbom/syft/syft/pkg/cataloger/internal/dependency"
)

// NewDBCataloger returns a new cataloger object initialized for arch linux pacman database flat-file stores.
func NewDBCataloger() pkg.Cataloger {
	return generic.NewCataloger("alpm-db-cataloger").
		WithParserByGlobs(parseAlpmDB, pkg.AlpmDBGlob).
		WithProcessors(dependency.Processor(dbEntryDependencySpecifier))
}
