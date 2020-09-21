package cfnpatcher

import (
	"context"

	"github.com/Jeffail/gabs/v2"
	"github.com/rs/zerolog/log"
)

type Configuration struct {
	Kilt string
	OptIn bool
}

const KiltIgnoreTag = "kiltignore"
const KiltIncludeTag = "kiltinclude"


func isIgnored(resource *gabs.Container, isOptIn bool) bool{
	tags := getTags(resource)
	_, included := tags[KiltIncludeTag]
	_, ignored := tags[KiltIgnoreTag]

	return !((isOptIn && included) || (!isOptIn && !ignored))
}

func Patch(ctx context.Context, config *Configuration , fragment []byte) ([]byte, error) {
	l := log.Ctx(ctx)
	template, err := gabs.ParseJSON(fragment)
	if err != nil {
		l.Error().Err(err).Msg("failed to parse input fragment")
		return nil, err
	}



	for name, resource := range template.S("Resources").ChildrenMap() {
		if matchFargate(resource) {
			if isIgnored(resource, config.OptIn) {
				l.Info().Str("resource", name).Msg("ignored resource due to tag")
				continue
			}
			l.Info().Str("resource", name).Msg("patching task definition")
			_, err := applyTaskDefinitionPatch(ctx, name, resource, config)
			if err != nil {
				l.Error().Err(err).Str("resource", name).Msgf("could not patch resource")
			}
		}
	}

	return template.Bytes(), nil
}
