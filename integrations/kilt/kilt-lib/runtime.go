package kilt_lib

import (
	"fmt"
	"github.com/go-akka/configuration"
)

func extractRuntime(config *configuration.Config) (*KiltRuntime, error) {
	r := new(KiltRuntime)

	if config.IsArray("runtime.upload") {
		uploads := config.GetValue("runtime.upload").GetArray()

		for k, u := range uploads {
			if u.IsObject() {
				var err error
				upload := u.GetObject()

				newUpload := new(KiltRuntimeUpload)

				newUpload.Payload, err = retrievePayload(upload)

				if err != nil {
					return nil, fmt.Errorf("could not extract payload for entry %d: %w", k, err)
				}

				newUpload.Destination = upload.GetKey("as").GetString()

				if newUpload.Destination == "" {
					return nil, fmt.Errorf("could not extract destination for entry %d: 'as' cannot be empty")
				}

				newUpload.Uid = getWithDefaultUint16(upload, "uid", defaultUserID)
				newUpload.Gid = getWithDefaultUint16(upload, "gid", defaultGroupID)
				newUpload.Permissions = getWithDefaultUint32(upload, "permissions", defaultPermissions)

				r.Uploads = append(r.Uploads, *newUpload)
			}
		}
	}

	if config.IsArray("runtime.exec") {
		for k, e := range config.GetValue("runtime.exec").GetArray() {
			if e.IsObject() {
				exec := e.GetObject()

				execParams := exec.GetKey("run").GetStringList()

				if len(execParams) == 0 {
					return nil, fmt.Errorf("could not add exec at entry %d: run cannot have 0 arguments", k)
				}

				r.Executables = append(r.Executables, KiltRuntimeExecutable{execParams})
			}
		}
	}

	return r, nil
}
