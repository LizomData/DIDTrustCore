package internal

import (
	"strings"

	"DIDTrustCore/util/sbom/syft/syft/artifact"
)

func ArtifactIDFromDigest(input string) artifact.ID {
	return artifact.ID(strings.TrimPrefix(input, "sha256:"))
}
