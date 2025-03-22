package matcher

import (
	"DIDTrustCore/util/grype/grype/match"
	"DIDTrustCore/util/grype/grype/matcher/apk"
	"DIDTrustCore/util/grype/grype/matcher/dotnet"
	"DIDTrustCore/util/grype/grype/matcher/dpkg"
	"DIDTrustCore/util/grype/grype/matcher/golang"
	"DIDTrustCore/util/grype/grype/matcher/java"
	"DIDTrustCore/util/grype/grype/matcher/javascript"
	"DIDTrustCore/util/grype/grype/matcher/msrc"
	"DIDTrustCore/util/grype/grype/matcher/portage"
	"DIDTrustCore/util/grype/grype/matcher/python"
	"DIDTrustCore/util/grype/grype/matcher/rpm"
	"DIDTrustCore/util/grype/grype/matcher/ruby"
	"DIDTrustCore/util/grype/grype/matcher/rust"
	"DIDTrustCore/util/grype/grype/matcher/stock"
)

// Config contains values used by individual matcher structs for advanced configuration
type Config struct {
	Java       java.MatcherConfig
	Ruby       ruby.MatcherConfig
	Python     python.MatcherConfig
	Dotnet     dotnet.MatcherConfig
	Javascript javascript.MatcherConfig
	Golang     golang.MatcherConfig
	Rust       rust.MatcherConfig
	Stock      stock.MatcherConfig
}

func NewDefaultMatchers(mc Config) []match.Matcher {
	return []match.Matcher{
		&dpkg.Matcher{},
		ruby.NewRubyMatcher(mc.Ruby),
		python.NewPythonMatcher(mc.Python),
		dotnet.NewDotnetMatcher(mc.Dotnet),
		&rpm.Matcher{},
		java.NewJavaMatcher(mc.Java),
		javascript.NewJavascriptMatcher(mc.Javascript),
		&apk.Matcher{},
		golang.NewGolangMatcher(mc.Golang),
		&msrc.Matcher{},
		&portage.Matcher{},
		rust.NewRustMatcher(mc.Rust),
		stock.NewStockMatcher(mc.Stock),
	}
}
