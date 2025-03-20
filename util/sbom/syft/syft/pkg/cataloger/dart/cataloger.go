/*
Package dart provides a concrete Cataloger implementations for the Dart language ecosystem.
*/
package dart

import (
	"DIDTrustCore/util/sbom/syft/syft/pkg"
	"DIDTrustCore/util/sbom/syft/syft/pkg/cataloger/generic"
)

// NewPubspecLockCataloger returns a new Dartlang cataloger object base on pubspec lock files.
func NewPubspecLockCataloger() pkg.Cataloger {
	return generic.NewCataloger("dart-pubspec-lock-cataloger").
		WithParserByGlobs(parsePubspecLock, "**/pubspec.lock")
}
