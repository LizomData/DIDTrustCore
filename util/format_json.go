package util

import (
	"DIDTrustCore/model"
	"encoding/json"
	"fmt"
)

func FormatJSON(data []byte) *model.SoftwareIdentity {
	var softwareIdentity model.SoftwareIdentity
	if err := json.Unmarshal(data, &softwareIdentity); err != nil {
		panic(fmt.Errorf("failed to parse JSON: %w", err))
	}
	return &softwareIdentity
}
