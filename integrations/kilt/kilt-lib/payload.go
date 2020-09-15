package kilt_lib

import (
	"encoding/base64"
	"fmt"
	"github.com/go-akka/configuration/hocon"
	"io/ioutil"
	"net/http"
)

func retrievePayload(object *hocon.HoconObject) ([]byte, error) {
	if object.GetKey("url").IsString() {
		return retrievePayloadViaURL(object.GetKey("url").GetString())
	} else if object.GetKey("file").IsString() {
		return retrievePayloadLocal(object.GetKey("file").GetString())
	} else if object.GetKey("payload").IsString() {
		return retrievePayloadBase64(object.GetKey("payload").GetString())
	} else if object.GetKey("text").IsString() {
		return []byte(object.GetKey("text").GetString()), nil
	}
	return nil, fmt.Errorf("could not retrieve payload")
}

func retrievePayloadViaURL(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	payload, err := ioutil.ReadAll(resp.Body)
	return payload, err
}

func retrievePayloadLocal(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func retrievePayloadBase64(encoded string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encoded)
}
