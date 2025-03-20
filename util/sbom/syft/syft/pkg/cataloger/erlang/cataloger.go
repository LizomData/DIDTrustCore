/*
Package erlang provides concrete Catalogers implementation relating to packages within the Erlang language ecosystem.
*/
package erlang

import (
	"DIDTrustCore/util/sbom/syft/syft/pkg"
	"DIDTrustCore/util/sbom/syft/syft/pkg/cataloger/generic"
)

// NewRebarLockCataloger returns a new cataloger instance for Erlang rebar.lock files.
func NewRebarLockCataloger() pkg.Cataloger {
	return generic.NewCataloger("erlang-rebar-lock-cataloger").
		WithParserByGlobs(parseRebarLock, "**/rebar.lock")
}

func NewOTPCataloger() pkg.Cataloger {
	return generic.NewCataloger("erlang-otp-application-cataloger").
		WithParserByGlobs(parseOTPApp, "**/*.app")
}
