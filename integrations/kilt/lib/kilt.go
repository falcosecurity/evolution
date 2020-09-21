package lib

import (
	"encoding/json"
	"fmt"
	"github.com/go-akka/configuration"
	"io/ioutil"
)

var defaults = `
build {
	entry_point: ${original.entry_point}
	command: ${original.command}
	image: ${original.image}
	environment_variables: ${original.environment_variables}

	mount: []
}
`

type Kilt struct {
	config *configuration.Config
	hasRuntimePatching bool
	// TODO(admiral0): add custom payload interpreter
}

func (k *Kilt) Build() (*KiltBuild, error) {
	return extractBuild(k.config)
}

func (k *Kilt) Runtime() (*KiltRuntime, error) {
	if k.hasRuntimePatching {
		return extractRuntime(k.config)
	}
	return nil, fmt.Errorf("kilt does not have 'runtime' instructions")
}

func (k *Kilt) HasRuntime() bool {
	return k.hasRuntimePatching
}

func KiltFromString(info *TargetInfo, config string) (*Kilt, error) {
	rawVars, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}

	initialConfig := "original:" + string(rawVars) + "\n" + defaults
	fullConfig := initialConfig + config

	userConfig := configuration.ParseString(fullConfig)

	kilt := &Kilt{
		userConfig,
		userConfig.HasPath("runtime"),
	}

	return kilt, nil
}

func KiltFromFile(info *TargetInfo, configurationPath string) (*Kilt, error) {
	config, err := ioutil.ReadFile(configurationPath)
	if err != nil {
		return nil, err
	}
	return KiltFromString(info, string(config))
}
