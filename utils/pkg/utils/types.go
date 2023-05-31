/*
Copyright (C) 2022 The Falco Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Maintainer struct {
	Name     string   `yaml:"name"`
	Github   string   `yaml:"github"`
	Company  string   `yaml:"company"`
	Projects []string `yaml:"projects"`
}

type Maintainers []Maintainer

type RepositoryScope string

type RepositoryStatus string

type Repository struct {
	Name        string           `yaml:"name"`
	Description string           `yaml:"description,omitempty"`
	Scope       RepositoryScope  `yaml:"scope"`
	Status      RepositoryStatus `yaml:"status,omitempty"`
}

type Repositories []Repository

const (
	RepositoryStatusStable     RepositoryStatus = "Stable"
	RepositoryStatusIncubating RepositoryStatus = "Incubating"
	RepositoryStatusSandbox    RepositoryStatus = "Sandbox"
	RepositoryStatusDeprecated RepositoryStatus = "Deprecated"
)

const (
	RepositoryScopeCore      RepositoryScope = "Core"
	RepositoryScopeEcosystem RepositoryScope = "Ecosystem"
	RepositoryScopeInfra     RepositoryScope = "Infra"
	RepositoryScopeSpecial   RepositoryScope = "Special"
)

func (r RepositoryStatus) String() string {
	return string(r)
}

func (r RepositoryScope) String() string {
	return string(r)
}

func (r *Repository) URL() string {
	return fmt.Sprintf("https://github.com/falcosecurity/%s", r.Name)
}

func readFromFile(path string, out interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, out)
}

func ReadRepositoriesFromFile(path string) (Repositories, error) {
	var res Repositories
	if err := readFromFile(path, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func ReadMaintainersFromFile(path string) (Maintainers, error) {
	var res Maintainers
	if err := readFromFile(path, &res); err != nil {
		return nil, err
	}
	return res, nil
}
