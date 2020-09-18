package lib

import (
	"encoding/json"
	"errors"
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

func KiltFromString(info *TargetInfo, config string) (*KiltBuild, *KiltRuntime, error) {
	rawVars, err := json.Marshal(info)
	if err != nil {
		return nil, nil, err
	}

	initialConfig := "original:" + string(rawVars) + "\n" + defaults
	fullConfig := initialConfig + config

	userConfig := configuration.ParseString(fullConfig)

	if !userConfig.HasPath("build") {
		return nil, nil, errors.New("kilt definition needs at least 'build' directive")
	}
	build, err := extractBuild(userConfig)

	if err != nil {
		return nil, nil, err
	}

	// runtime patching is optional
	if !userConfig.HasPath("runtime") {
		return build, nil, nil
	}

	run, err := extractRuntime(userConfig)

	if err != nil {
		return nil, nil, err
	}

	return build, run, nil
}

func KiltFromFile(info *TargetInfo, configurationPath string) (*KiltBuild, *KiltRuntime, error) {
	config, err := ioutil.ReadFile(configurationPath)
	if err != nil {
		return nil, nil, err
	}
	return KiltFromString(info, string(config))
}
