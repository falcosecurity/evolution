package kilt_lib

import (
	"errors"
	"github.com/go-akka/configuration"
	"io/ioutil"
)

type defaultData struct {
	Original TargetInfo `json:"original"`
}

var defaults = configuration.ParseString(`
build {
	entry_point: ${original.entry_point}
	command: ${original.command}
	image: ${original.image}
	environment_variables: ${original.environment_variables}

	mount: []
}
`)

func KiltFromString(info TargetInfo, config string) (*KiltBuild, *KiltRuntime, error) {
	vars := configuration.FromObject(defaultData{info})
	kiltDefaults := configuration.NewConfigFromConfig(defaults, vars)

	userConfig := configuration.ParseString(config)

	finalConfig := configuration.NewConfigFromConfig(userConfig, kiltDefaults)

	if !userConfig.HasPath("build") {
		return nil, nil, errors.New("kilt definition needs at least 'build' directive")
	}
	build, err := extractBuild(finalConfig)

	if err != nil {
		return nil, nil, err
	}

	// runtime patching is optional
	if !userConfig.HasPath("runtime") {
		return build, nil, nil
	}

	run, err := extractRuntime(finalConfig)

	if err != nil {
		return nil, nil, err
	}

	return build, run, nil
}

func KiltFromFile(info TargetInfo, configurationPath string) (*KiltBuild, *KiltRuntime, error) {
	config, err := ioutil.ReadFile(configurationPath)
	if err != nil {
		return nil, nil, err
	}
	return KiltFromString(info, string(config))
}
