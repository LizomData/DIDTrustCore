package main

import (
	"fmt"
	"os"

	"DIDTrustCore/util/sbom/syft/syft/pkg/cataloger/binary/test-fixtures/manager/internal/cli"
	"DIDTrustCore/util/sbom/syft/syft/pkg/cataloger/binary/test-fixtures/manager/internal/ui"
)

func main() {
	cmd, err := cli.New()
	if err != nil {
		exit(err)
	}

	if err := cmd.Execute(); err != nil {
		exit(err)
	}
}

func exit(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", ui.RenderError(err))
	}
	os.Exit(1)
}
