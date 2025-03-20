package wordpress

import (
	"DIDTrustCore/util/sbom/syft/syft/pkg"
	"DIDTrustCore/util/sbom/syft/syft/pkg/cataloger/generic"
)

const (
	catalogerName        = "wordpress-plugins-cataloger"
	wordpressPluginsGlob = "**/wp-content/plugins/*/*.php"
)

func NewWordpressPluginCataloger() pkg.Cataloger {
	return generic.NewCataloger(catalogerName).
		WithParserByGlobs(parseWordpressPluginFiles, wordpressPluginsGlob)
}
