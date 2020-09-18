package lib

import (
	"github.com/go-akka/configuration"
	"github.com/go-akka/configuration/hocon"
)

func extractToStringMap(config *configuration.Config, path string) map[string]string {
	value := make(map[string]string)

	if config.HasPath(path) && config.IsObject(path) {
		obj := config.GetNode(path).GetObject()

		for k, v := range obj.Items() {
			value[k] = v.GetString()
		}
	}

	return value
}

func getWithDefaultUint16(object *hocon.HoconObject, key string, fallback uint16) uint16 {
	t := object.GetKey(key)

	if t == nil || t.IsEmpty() {
		return fallback
	}

	return uint16(t.GetInt32())
}

func getWithDefaultUint32(object *hocon.HoconObject, key string, fallback uint32) uint32 {
	t := object.GetKey(key)

	if  t==nil ||  t.IsEmpty() {
		return fallback
	}

	return uint32(t.GetInt64())
}


func emptyIncludeCallback(filename string) *hocon.HoconRoot {
	return hocon.Parse("", emptyIncludeCallback)
}