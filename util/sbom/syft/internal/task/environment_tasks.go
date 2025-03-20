package task

import (
	"context"

	"DIDTrustCore/util/sbom/syft/internal/sbomsync"
	"DIDTrustCore/util/sbom/syft/syft/file"
	"DIDTrustCore/util/sbom/syft/syft/linux"
)

// TODO: add tui element here?

func NewEnvironmentTask() Task {
	fn := func(_ context.Context, resolver file.Resolver, builder sbomsync.Builder) error {
		release := linux.IdentifyRelease(resolver)
		if release != nil {
			builder.SetLinuxDistribution(*release)
		}

		return nil
	}

	return NewTask("environment-cataloger", fn)
}
