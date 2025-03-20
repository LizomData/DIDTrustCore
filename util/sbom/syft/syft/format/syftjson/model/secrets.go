package model

import (
	"DIDTrustCore/util/sbom/syft/syft/file"
)

type Secrets struct {
	Location file.Coordinates    `json:"location"`
	Secrets  []file.SearchResult `json:"secrets"`
}
