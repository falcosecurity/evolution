package cfnpatcher

import (
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"os"
)

func getTags(template *gabs.Container) map[string]string {
	tags := make(map[string]string)
	if template.Exists("Properties", "Tags") {
		for _, tag := range template.S("Properties", "Tags").Children() {
			if tag.Exists("Key") && tag.Exists("Value"){
				tags[tag.S("Key").Data().(string)] = tag.S("Value").Data().(string)
			}
		}
	}
	return tags
}

func exitErrorf(msg string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
