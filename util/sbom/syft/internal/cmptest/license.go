package cmptest

import (
	"github.com/google/go-cmp/cmp"

	"DIDTrustCore/util/sbom/syft/syft/file"
	"DIDTrustCore/util/sbom/syft/syft/pkg"
)

type LicenseComparer func(x, y pkg.License) bool

func DefaultLicenseComparer(x, y pkg.License) bool {
	return cmp.Equal(
		x, y,
		cmp.Comparer(DefaultLocationComparer),
		cmp.Comparer(buildSetComparer[file.Location, file.LocationSet](DefaultLocationComparer)),
	)
}

func LicenseComparerWithoutLocationLayer(x, y pkg.License) bool {
	return cmp.Equal(
		x, y,
		cmp.Comparer(LocationComparerWithoutLayer),
		cmp.Comparer(buildSetComparer[file.Location, file.LocationSet](LocationComparerWithoutLayer)),
	)
}
