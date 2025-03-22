package main

import (
	"DIDTrustCore/util/grype/cmd/grype/cli"
	"DIDTrustCore/util/grype/cmd/grype/cli/commands"
	"DIDTrustCore/util/grype/cmd/grype/cli/options"
	"fmt"
	"github.com/anchore/clio"
)

// applicationName is the non-capitalized name of the application (do not change this)
const applicationName = "grype"

// all variables here are provided as build-time arguments, with clear default values
var (
	version        = "[not provided]"
	buildDate      = "[not provided]"
	gitCommit      = "[not provided]"
	gitDescription = "[not provided]"
)

func main() {
	app := cli.Application(
		clio.Identification{
			Name:           applicationName,
			Version:        version,
			BuildDate:      buildDate,
			GitCommit:      gitCommit,
			GitDescription: gitDescription,
		},
	)

	opts := options.DefaultGrype(app.ID())
	opts.Outputs = []string{"json"}
	opts.File = "scanResult/grype-report.json"
	opts.Pretty = true

	userInput := "sbom:/Users/q/Downloads/bom.spdx.json"
	if err := runGrypeWrapper(app, opts, userInput); err != nil {
		panic(err)
	}
}

func runGrypeWrapper(app clio.Application, opts *options.Grype, userInput string) error {

	if err := commands.RunGrype(app, opts, userInput); err != nil {
		return fmt.Errorf("扫描失败: %w", err)
	}

	return nil
}
