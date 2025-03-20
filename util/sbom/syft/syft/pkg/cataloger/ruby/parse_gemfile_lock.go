package ruby

import (
	"bufio"
	"context"
	"strings"

	"github.com/scylladb/go-set/strset"

	"DIDTrustCore/util/sbom/syft/internal/unknown"
	"DIDTrustCore/util/sbom/syft/syft/artifact"
	"DIDTrustCore/util/sbom/syft/syft/file"
	"DIDTrustCore/util/sbom/syft/syft/pkg"
	"DIDTrustCore/util/sbom/syft/syft/pkg/cataloger/generic"
)

var _ generic.Parser = parseGemFileLockEntries

var sectionsOfInterest = strset.New("GEM", "GIT", "PATH", "PLUGIN SOURCE")

// parseGemFileLockEntries is a parser function for Gemfile.lock contents, returning all Gems discovered.
func parseGemFileLockEntries(_ context.Context, _ file.Resolver, _ *generic.Environment, reader file.LocationReadCloser) ([]pkg.Package, []artifact.Relationship, error) {
	var pkgs []pkg.Package
	scanner := bufio.NewScanner(reader)

	var currentSection string

	for scanner.Scan() {
		line := scanner.Text()
		sanitizedLine := strings.TrimSpace(line)

		if len(line) > 1 && line[0] != ' ' {
			// start of section
			currentSection = sanitizedLine
			continue
		} else if !sectionsOfInterest.Has(currentSection) {
			// skip this line, we're in the wrong section
			continue
		}

		if isDependencyLine(line) {
			candidate := strings.Fields(sanitizedLine)
			if len(candidate) != 2 {
				continue
			}
			pkgs = append(pkgs,
				newGemfileLockPackage(
					candidate[0],
					strings.Trim(candidate[1], "()"),
					reader.Location.WithAnnotation(pkg.EvidenceAnnotationKey, pkg.PrimaryEvidenceAnnotation),
				),
			)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}
	return pkgs, nil, unknown.IfEmptyf(pkgs, "unable to determine packages")
}

func isDependencyLine(line string) bool {
	if len(line) < 5 {
		return false
	}
	return strings.Count(line[:5], " ") == 4
}
