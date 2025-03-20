/*
Package cpp provides a concrete Cataloger implementations for the C/C++ language ecosystem.
*/
package cpp

import (
	"DIDTrustCore/util/sbom/syft/syft/pkg"
	"DIDTrustCore/util/sbom/syft/syft/pkg/cataloger/generic"
)

// NewConanCataloger returns a new C/C++ conanfile.txt and conan.lock cataloger object.
func NewConanCataloger() pkg.Cataloger {
	return generic.NewCataloger("conan-cataloger").
		WithParserByGlobs(parseConanfile, "**/conanfile.txt").
		WithParserByGlobs(parseConanLock, "**/conan.lock")
}

// NewConanInfoCataloger returns a new C/C++ conaninfo.txt cataloger object.
func NewConanInfoCataloger() pkg.Cataloger {
	return generic.NewCataloger("conan-info-cataloger").
		WithParserByGlobs(parseConaninfo, "**/conaninfo.txt")
}
