package grypeService

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

type GrypeServiceConfig struct {
	Opts       *options.Grype
	OutputPath string
	PublicPath string
}
type GrypeService struct {
	App    clio.Application
	Config *GrypeServiceConfig
}

var Service = initService()

func initService() *GrypeService {
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

	config := GrypeServiceConfig{opts, "./tmp/scanResult/", "/scanResult/"}

	//userInput := "sbom:/Users/q/Downloads/bom.spdx.json"
	//if err := runGrypeWrapper(app, opts, userInput); err != nil {
	//	panic(err)
	//}
	return &GrypeService{app, &config}
}

func (u *GrypeService) RunGrypeWrapper(outputName, sbomPath string) error {
	u.Config.Opts.File = u.Config.OutputPath + outputName
	userInput := "sbom:" + sbomPath
	if err := commands.RunGrype(u.App, u.Config.Opts, userInput); err != nil {
		return fmt.Errorf("扫描失败: %w", err)
	}
	return nil
}
