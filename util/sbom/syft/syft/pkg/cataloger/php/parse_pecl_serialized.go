package php

import (
	"context"
	"fmt"
	"io"

	"github.com/elliotchance/phpserialize"

	"DIDTrustCore/util/sbom/syft/internal/log"
	"DIDTrustCore/util/sbom/syft/syft/artifact"
	"DIDTrustCore/util/sbom/syft/syft/file"
	"DIDTrustCore/util/sbom/syft/syft/pkg"
	"DIDTrustCore/util/sbom/syft/syft/pkg/cataloger/generic"
)

// parsePeclSerialized is a parser function for PECL metadata contents, returning "Default" php packages discovered.
func parsePeclSerialized(_ context.Context, _ file.Resolver, _ *generic.Environment, reader file.LocationReadCloser) ([]pkg.Package, []artifact.Relationship, error) {
	var pkgs []pkg.Package
	data, err := io.ReadAll(reader)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to read file: %w", err)
	}

	metadata, err := phpserialize.UnmarshalAssociativeArray(
		data,
	)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse pecl metadata file: %w", err)
	}

	name, ok := metadata["name"].(string)
	if !ok {
		return nil, nil, fmt.Errorf("failed to parse pecl package name: %w", err)
	}

	version := readStruct(metadata, "version", "release")
	license := readStruct(metadata, "license", "_content")

	pkgs = append(
		pkgs,
		newPeclPackage(
			pkg.PhpPeclEntry{
				Name:    name,
				Version: version,
				License: []string{
					license,
				},
			},
			reader.Location,
		),
	)

	return pkgs, nil, nil
}

func readStruct(metadata any, fields ...string) string {
	if len(fields) > 0 {
		value, ok := metadata.(map[any]any)
		if !ok {
			log.Tracef("unable to read '%s' from: %v", fields[0], metadata)
			return ""
		}
		return readStruct(value[fields[0]], fields[1:]...)
	}
	value, ok := metadata.(string)
	if !ok {
		log.Tracef("unable to read value from: %v", metadata)
	}
	return value
}
