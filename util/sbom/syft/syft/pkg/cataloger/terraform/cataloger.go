package terraform

import (
	"DIDTrustCore/util/sbom/syft/syft/pkg"
	"DIDTrustCore/util/sbom/syft/syft/pkg/cataloger/generic"
)

func NewLockCataloger() pkg.Cataloger {
	return generic.NewCataloger("terraform-lock-cataloger").
		WithParserByGlobs(parseTerraformLock, "**/.terraform.lock.hcl")
}
