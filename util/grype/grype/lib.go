package grype

import (
	"github.com/wagoodman/go-partybus"

	"github.com/anchore/go-logger"
	"DIDTrustCore/util/grype/internal/bus"
	"DIDTrustCore/util/grype/internal/log"
)

func SetLogger(l logger.Logger) {
	log.Set(l)
}

func SetBus(b *partybus.Bus) {
	bus.Set(b)
}
