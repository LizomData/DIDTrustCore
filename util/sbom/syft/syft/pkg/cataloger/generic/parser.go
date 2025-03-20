package generic

import (
	"context"

	"DIDTrustCore/util/sbom/syft/syft/artifact"
	"DIDTrustCore/util/sbom/syft/syft/file"
	"DIDTrustCore/util/sbom/syft/syft/linux"
	"DIDTrustCore/util/sbom/syft/syft/pkg"
)

type Environment struct {
	LinuxRelease *linux.Release
}

type Parser func(context.Context, file.Resolver, *Environment, file.LocationReadCloser) ([]pkg.Package, []artifact.Relationship, error)
