package lib

const (
	defaultUserID      = 0
	defaultGroupID     = 0
	defaultPermissions = 0755
)

type TargetInfo struct {
	Image                string            `json:"image"`
	ContainerName        string            `json:"container_name"`
	ContainerGroupName   string            `json:"container_group_name"`
	EntryPoint           []string          `json:"entry_point"`
	Command              []string          `json:"command"`
	EnvironmentVariables map[string]string `json:"environment_variables"`
	Metadata             map[string]string `json:"metadata"`
}

type KiltBuildResource struct {
	Name string
	Image      string
	Volumes    []string
	EntryPoint []string
}

type KiltBuild struct {
	Image string
	EntryPoint           []string
	Command              []string
	EnvironmentVariables map[string]string

	Resources []KiltBuildResource
}

type KiltRuntimeUpload struct {
	Payload     []byte
	Destination string
	Uid         uint16
	Gid         uint16
	Permissions uint32
}

type KiltRuntimeExecutable struct {
	Run []string
}

type KiltRuntime struct {
	Uploads     []KiltRuntimeUpload
	Executables []KiltRuntimeExecutable
}
