package cfnpatcher

import (
	"github.com/Jeffail/gabs/v2"
	"github.com/admiral0/evolution/integrations/kilt/lib"
)

func extractInfo(template *gabs.Container) lib.TargetInfo
