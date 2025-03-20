package task

import (
	"DIDTrustCore/util/sbom/syft/syft/cataloging"
	"DIDTrustCore/util/sbom/syft/syft/cataloging/filecataloging"
	"DIDTrustCore/util/sbom/syft/syft/cataloging/pkgcataloging"
)

type CatalogingFactoryConfig struct {
	ComplianceConfig     cataloging.ComplianceConfig
	SearchConfig         cataloging.SearchConfig
	RelationshipsConfig  cataloging.RelationshipsConfig
	DataGenerationConfig cataloging.DataGenerationConfig
	LicenseConfig        cataloging.LicenseConfig
	PackagesConfig       pkgcataloging.Config
	FilesConfig          filecataloging.Config
}

func DefaultCatalogingFactoryConfig() CatalogingFactoryConfig {
	return CatalogingFactoryConfig{
		ComplianceConfig:     cataloging.DefaultComplianceConfig(),
		SearchConfig:         cataloging.DefaultSearchConfig(),
		RelationshipsConfig:  cataloging.DefaultRelationshipsConfig(),
		DataGenerationConfig: cataloging.DefaultDataGenerationConfig(),
		LicenseConfig:        cataloging.DefaultLicenseConfig(),
		PackagesConfig:       pkgcataloging.DefaultConfig(),
		FilesConfig:          filecataloging.DefaultConfig(),
	}
}
