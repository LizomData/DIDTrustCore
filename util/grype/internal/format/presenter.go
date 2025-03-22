package format

import (
	"github.com/wagoodman/go-presenter"

	"DIDTrustCore/util/grype/grype/presenter/cyclonedx"
	"DIDTrustCore/util/grype/grype/presenter/json"
	"DIDTrustCore/util/grype/grype/presenter/models"
	"DIDTrustCore/util/grype/grype/presenter/sarif"
	"DIDTrustCore/util/grype/grype/presenter/table"
	"DIDTrustCore/util/grype/grype/presenter/template"
	"DIDTrustCore/util/grype/internal/log"
)

type PresentationConfig struct {
	TemplateFilePath string
	ShowSuppressed   bool
	Pretty           bool
}

// GetPresenter retrieves a Presenter that matches a CLI option
func GetPresenter(format Format, c PresentationConfig, pb models.PresenterConfig) presenter.Presenter {
	switch format {
	case JSONFormat:
		return json.NewPresenter(pb)
	case TableFormat:
		return table.NewPresenter(pb, c.ShowSuppressed)

	// NOTE: cyclonedx is identical to EmbeddedVEXJSON
	// The cyclonedx library only provides two BOM formats: JSON and XML
	// These embedded formats will be removed in v1.0
	case CycloneDXFormat:
		return cyclonedx.NewXMLPresenter(pb)
	case CycloneDXJSON:
		return cyclonedx.NewJSONPresenter(pb)
	case CycloneDXXML:
		return cyclonedx.NewXMLPresenter(pb)
	case SarifFormat:
		return sarif.NewPresenter(pb)
	case TemplateFormat:
		return template.NewPresenter(pb, c.TemplateFilePath)
	// DEPRECATED TODO: remove in v1.0
	case EmbeddedVEXJSON:
		log.Warn("embedded-cyclonedx-vex-json format is deprecated and will be removed in v1.0")
		return cyclonedx.NewJSONPresenter(pb)
	case EmbeddedVEXXML:
		log.Warn("embedded-cyclonedx-vex-xml format is deprecated and will be removed in v1.0")
		return cyclonedx.NewXMLPresenter(pb)
	default:
		return nil
	}
}
