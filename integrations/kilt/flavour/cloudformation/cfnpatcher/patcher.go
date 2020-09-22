package cfnpatcher

import (
	"context"
	"fmt"
	"github.com/admiral0/evolution/integrations/kilt/lib"

	"github.com/Jeffail/gabs/v2"
	"github.com/rs/zerolog/log"
)

const KiltImageName = "KiltImage"

func applyTaskDefinitionPatch(ctx context.Context, name string, resource *gabs.Container, configuration *Configuration) (*gabs.Container, error) {
	l := log.Ctx(ctx)
	successes := 0
	containers := make(map[string]*lib.KiltBuildResource)
	if resource.Exists("Properties", "ContainerDefinitions") {
		for _, container := range resource.S("Properties", "ContainerDefinitions").Children() {
			info := extractContainerInfo(resource, name, container)
			kilt, err := lib.KiltFromString(info, configuration.Kilt)
			if err != nil {
				return nil, fmt.Errorf("could not construct kilt: %w", err)
			}
			patch, err  := kilt.Build()
			if err != nil {
				return nil, fmt.Errorf("could not construct kilt patch: %w", err)
			}
			err = applyContainerDefinitionPatch(l.WithContext(ctx), container, patch)
			if err != nil {
				l.Warn().Str("resource", name).Err(err).Msg("skipped patching container in task definition")
			}else{
				successes += 1
			}

			for _, appendResource := range patch.Resources {
				containers[appendResource.Name] = &appendResource
			}
		}
		err := appendContainers(resource, containers)
		if err != nil {
			return nil, err
		}
	}
	if successes == 0 {
		return resource, fmt.Errorf("could not patch a single container in the task")
	}
	return resource, nil
}

func applyContainerDefinitionPatch(ctx context.Context, container *gabs.Container, patch *lib.KiltBuild) error {
	_, err := container.Set(patch.EntryPoint, "EntryPoint")
	if err != nil {
		return fmt.Errorf("could not set EntryPoint: %w", err)
	}
	_, err = container.Set(patch.Command, "Command")
	if err != nil {
		return fmt.Errorf("could not set Command: %w", err)
	}
	_, err = container.Set(patch.Image, "Image")
	if err != nil {
		return fmt.Errorf("could not set Command: %w", err)
	}

	if ! container.Exists("VolumesFrom") {
		_, err = container.Set([]interface{}{}, "VolumesFrom")
		if err!= nil {
			return fmt.Errorf("could not set VolumesFrom: %w", err)
		}
	}

	for _, newContainer := range patch.Resources {
		addVolume := map[string]interface{}{
			"ReadOnly": true,
			"SourceContainer": newContainer.Name,
		}

		_, err = container.Set(addVolume, "VolumesFrom", "-")
		if err != nil {
			return fmt.Errorf("could not add VolumesFrom directive: %w", err)
		}
	}


	if ! (container.Exists("Environment") || len(patch.EnvironmentVariables) == 0) {
		_, err = container.Set([]interface{}{}, "Environment")

		if err != nil {
			return fmt.Errorf("could not add environment variable container: %w", err)
		}
	}
	for k, v := range patch.EnvironmentVariables {
		keyval := make(map[string]string)
		keyval["Name"] = k
		keyval["Value"] = v
		_, err = container.Set(keyval, "Environment", "-")

		if err != nil {
			return fmt.Errorf("could not add environment variable %v: %w", keyval, err)
		}

	}

	return nil
}

func appendContainers(resource *gabs.Container, containers map[string]*lib.KiltBuildResource) error {
	for _, inject := range containers {
		_, err := resource.Set(map[string]interface{}{
			"Name": inject.Name,
			"Image": inject.Image,
			"EntryPoint": inject.EntryPoint,
		}, "Properties", "ContainerDefinitions", "-")
		if err != nil {
			return fmt.Errorf("could not inject %s: %w", inject.Name, err)
		}
	}
	return nil
}