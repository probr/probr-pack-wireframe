package pack

import (
	"github.com/markbates/pkger"
	"github.com/probr/probr-pack-wireframe/internal/welcome"
	"github.com/probr/probr-sdk/probeengine"
)

// GetProbes returns a list of probe objects
func GetProbes() []probeengine.Probe {
	return []probeengine.Probe{
		welcome.Probe,
	}
}

func init() {
	// pkger.Include is a no-op that directs the pkger tool to include the desired file or folder.
	pkger.Include("/internal/welcome/welcome.feature")
}
