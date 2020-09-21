package cfnpatcher

import (
	"context"
	"fmt"

	"github.com/Jeffail/gabs/v2"
	"github.com/rs/zerolog/log"
)

const KiltImageName = "KiltImage"

func applyTaskDefinitionPatch(ctx context.Context, name string, resource *gabs.Container, config Configuration) (*gabs.Container, error) {
	l := log.Ctx(ctx)
	successes := 0
	if resource.Exists("Properties", "ContainerDefinitions") {
		for _, container := range resource.S("Properties", "ContainerDefinitions").Children() {
			err := applyContainerDefinitionPatch(l.WithContext(ctx), container, config)
			if err != nil {
				l.Warn().Str("resource", name).Err(err).Msg("skipped patching container in task definition")
			}else{
				successes += 1
			}
		}
		err := appendKiltContainer(resource, config)
		if err != nil {
			return nil, err
		}
	}
	if successes == 0 {
		return resource, fmt.Errorf("could not patch a single container in the task")
	}
	return resource, nil
}

func applyContainerDefinitionPatch(ctx context.Context, container *gabs.Container, config Configuration) error {
	cmdline := make([]string, 0)

	if container.Exists("EntryPoint") {
		for _, arg := range container.S("EntryPoint").Children() {
			cmdline = append(cmdline, arg.Data().(string))
		}
	}
	if container.Exists("Command") {
		for _, arg := range container.S("Command").Children() {
			cmdline = append(cmdline, arg.Data().(string))
		}
	}

	if len(cmdline) == 0 {
		return fmt.Errorf("container has no specified entrypoint or command")
	}


	_, err := container.Set([]string{config.Executable, "--"}, "EntryPoint")
	if err != nil {
		return fmt.Errorf("could not set EntryPoint: %w", err)
	}
	_, err = container.Set(cmdline, "Command")
	if err != nil {
		return fmt.Errorf("could not set Command: %w", err)
	}

	volumesFrom := map[string]interface{}{
		"ReadOnly": true,
		"SourceContainer": KiltImageName,
	}

	if container.Exists("VolumesFrom") {
		_, err = container.Set(volumesFrom, "VolumesFrom", "-")
	}else{
		_, err = container.Set([]interface{}{volumesFrom}, "VolumesFrom")
	}

	if err!= nil {
		return fmt.Errorf("could not set VolumesFrom: %w", err)
	}

	// TODO metadata as an environment variable (__KILT_METADATA)

	return nil
}

func appendKiltContainer(resource *gabs.Container, config Configuration) error {
	_, err := resource.Set(map[string]interface{}{
		"Name": KiltImageName,
		"Image": config.Image,
	}, "Properties", "ContainerDefinitions", "-")

	return err
}