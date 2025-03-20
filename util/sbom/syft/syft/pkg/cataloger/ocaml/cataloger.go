/*
Package ocaml provides a concrete Cataloger implementation for packages relating to the OCaml language ecosystem.
*/
package ocaml

import (
	"DIDTrustCore/util/sbom/syft/syft/pkg"
	"DIDTrustCore/util/sbom/syft/syft/pkg/cataloger/generic"
)

// NewOpamPackageManagerCataloger returns a new cataloger object for OCaml opam.
func NewOpamPackageManagerCataloger() pkg.Cataloger {
	return generic.NewCataloger("opam-cataloger").
		WithParserByGlobs(parseOpamPackage, "*opam")
}
