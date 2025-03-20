package githubactions

import (
	"context"
	"fmt"
	"io"

	"gopkg.in/yaml.v3"

	"DIDTrustCore/util/sbom/syft/internal/unknown"
	"DIDTrustCore/util/sbom/syft/syft/artifact"
	"DIDTrustCore/util/sbom/syft/syft/file"
	"DIDTrustCore/util/sbom/syft/syft/pkg"
	"DIDTrustCore/util/sbom/syft/syft/pkg/cataloger/generic"
)

var _ generic.Parser = parseCompositeActionForActionUsage

type compositeActionDef struct {
	Runs compositeActionRunsDef `yaml:"runs"`
}

type compositeActionRunsDef struct {
	Steps []stepDef `yaml:"steps"`
}

func parseCompositeActionForActionUsage(_ context.Context, _ file.Resolver, _ *generic.Environment, reader file.LocationReadCloser) ([]pkg.Package, []artifact.Relationship, error) {
	contents, errs := io.ReadAll(reader)
	if errs != nil {
		return nil, nil, fmt.Errorf("unable to read yaml composite action file: %w", errs)
	}

	var ca compositeActionDef
	if errs = yaml.Unmarshal(contents, &ca); errs != nil {
		return nil, nil, fmt.Errorf("unable to parse yaml composite action file: %w", errs)
	}

	// we use a collection to help with deduplication before raising to higher level processing
	pkgs := pkg.NewCollection()

	for _, step := range ca.Runs.Steps {
		if step.Uses == "" {
			continue
		}

		p, err := newPackageFromUsageStatement(step.Uses, reader.Location)
		if err != nil {
			errs = unknown.Append(errs, reader, err)
		}
		if p != nil {
			pkgs.Add(*p)
		}
	}

	return pkgs.Sorted(), nil, errs
}
