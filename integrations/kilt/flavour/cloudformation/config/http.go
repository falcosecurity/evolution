package config

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func FromWeb(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(fmt.Errorf("could not perform http request: %w", err))
	}
	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		panic(fmt.Errorf("could not read http response: %w", err))
	}

	return string(data)
}