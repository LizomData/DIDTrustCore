package main

import (
	"context"
	"fmt"
	"DIDTrustCore/util/sbom/syft/syft/format/cyclonedxjson"
	"DIDTrustCore/util/sbom/syft/syft/format/spdxjson"
	"os"

	"DIDTrustCore/util/sbom/syft/syft"
	"DIDTrustCore/util/sbom/syft/syft/format"
	"DIDTrustCore/util/sbom/syft/syft/format/syftjson"
	"DIDTrustCore/util/sbom/syft/syft/sbom"
	"DIDTrustCore/util/sbom/syft/syft/source"
)

const defaultImage = "alpine:3.19"

func main() {
	// automagically get a source.Source for arbitrary string input
	//src := getSource(imageReference())
	src := getSource("/Users/q/GolandProjects/syft/syft")

	// catalog the given source and return a SBOM
	sbom := getSBOM(src)

	// take the SBOM object and encode it into the syft-json representation
	bytes := formatSBOM_spdx(sbom)
	//bytes := formatSBOM_cdx(sbom)

	// show the SBOM!
	//fmt.Println(string(bytes))

	err := os.WriteFile("sbom.syft.json", bytes, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("SBOM 已保存至 sbom.syft.json")
}

func imageReference() string {
	// read an image string reference from the command line or use a default
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	return defaultImage
}

func getSource(input string) source.Source {
	src, err := syft.GetSource(context.Background(), input, nil)

	if err != nil {
		panic(err)
	}

	return src
}

func getSBOM(src source.Source) sbom.SBOM {
	s, err := syft.CreateSBOM(context.Background(), src, nil)
	if err != nil {
		panic(err)
	}

	return *s
}

func formatSBOM(s sbom.SBOM) []byte {
	bytes, err := format.Encode(s, syftjson.NewFormatEncoder())
	if err != nil {
		panic(err)
	}
	return bytes
}
func formatSBOM_spdx(s sbom.SBOM) []byte {
	e, _ := spdxjson.NewFormatEncoderWithConfig(spdxjson.DefaultEncoderConfig())
	bytes, err := format.Encode(s, e)
	if err != nil {
		panic(err)
	}
	return bytes
}

func formatSBOM_cdx(s sbom.SBOM) []byte {
	e, _ := cyclonedxjson.NewFormatEncoderWithConfig(cyclonedxjson.DefaultEncoderConfig())
	bytes, err := format.Encode(s, e)
	if err != nil {
		panic(err)
	}
	return bytes
}
