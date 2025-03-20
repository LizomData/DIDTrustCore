package cpe

import (
	"DIDTrustCore/util/sbom/syft/syft/cpe"
	"DIDTrustCore/util/sbom/syft/syft/pkg"
	"DIDTrustCore/util/sbom/syft/syft/pkg/cataloger/internal/cpegenerate"
)

func Generate(p pkg.Package) []cpe.CPE {
	return cpegenerate.FromPackageAttributes(p)
}

func DictionaryFind(p pkg.Package) ([]cpe.CPE, bool) {
	return cpegenerate.FromDictionaryFind(p)
}
