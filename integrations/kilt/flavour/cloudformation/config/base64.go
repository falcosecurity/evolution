package config

import (
	"encoding/base64"
	"fmt"
)

func FromBase64(payload string, decompress bool) string {
	var data []byte
	_, err := base64.StdEncoding.Decode(data, []byte(payload))

	if err != nil {
		panic(fmt.Errorf("could not decode base64: %w", err))
	}

	if decompress {
		data = decompressBytes(data)
	}

	return string(data)
}