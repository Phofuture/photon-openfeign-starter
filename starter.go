package openFiegnStarter

import (
	"github.com/Phofuture/photon-core-starter/core"
	"github.com/Phofuture/photon-openfeign-starter/openfeign"
)

func init() {
	core.RegisterCoreDependency(openfeign.StartOpenFeign)
}
