package cfnpatcher

import (
	"github.com/Jeffail/gabs/v2"
	"github.com/admiral0/evolution/integrations/kilt/lib"
)

func extractContainerInfo(group *gabs.Container, groupName string, container *gabs.Container) *lib.TargetInfo {
	info := new(lib.TargetInfo)

	info.ContainerName = container.S("Name").Data().(string)
	info.ContainerGroupName = groupName
	info.EnvironmentVariables = make(map[string]string)
	info.Metadata = make(map[string]string)

	if container.Exists("Image") {
		info.Image = container.S("Image").Data().(string)
	}

	if container.Exists("EntryPoint") {
		for _, arg := range container.S("EntryPoint").Children() {
			info.EntryPoint = append(info.EntryPoint, arg.Data().(string))
		}
	}

	if container.Exists("Command") {
		for _, arg := range container.S("Command").Children() {
			info.Command = append(info.Command, arg.Data().(string))
		}
	}

	if container.Exists("Environment") {
		for _, env := range container.S("Environment").Children() {
			info.EnvironmentVariables[env.S("Name").Data().(string)] = env.S("Value").Data().(string)
		}
	}

	if group.Exists("Properties", "Tags") {
		for _, tag := range group.S("Properties", "Tags").Children() {
			if tag.Exists("Key") && tag.Exists("Value") {
				info.Metadata[tag.S("Key").Data().(string)] = tag.S("Value").Data().(string)
			}
		}
	}

	// TODO(admiral0): metadata tags

	return info
}
