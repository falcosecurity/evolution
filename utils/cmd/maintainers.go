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

package main

import (
	"fmt"
	"strings"

	"github.com/falcosecurity/evolution/utils/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	maintainersReposFilePath  string
	maintainersInFilePath     string
	maintainersOutFilePath    string
	maintainersTextStartTag   = "<!-- MAINTAINERS-LIST -->\n"
	maintainersTextEndTag     = "<!-- /MAINTAINERS-LIST -->\n"
	maintainersTextTitleStyle = "##"
)

func maintainersTextEditor(s string) (string, error) {
	if len(s) == 0 {
		s = maintainersTextStartTag + maintainersTextEndTag
	}
	maintainers, err := utils.ReadMaintainersFromFile(maintainersInFilePath)
	if err != nil {
		return "", err
	}
	repositories, err := utils.ReadRepositoriesFromFile(maintainersReposFilePath)
	if err != nil {
		return "", err
	}

	var mList utils.Maintainers
	var mCoreList utils.Maintainers
	for _, m := range maintainers {
		listAdded := false
		listCoreAdded := false
		for _, r := range repositories {
			for _, url := range m.Projects {
				if url == r.URL() {
					if !listAdded {
						mList = append(mList, m)
						listAdded = true
					}
					if !listCoreAdded && r.Status == utils.RepositoryStatusOfficial {
						mCoreList = append(mCoreList, m)
						listCoreAdded = true
					}
				}
			}
		}
	}

	var res strings.Builder
	res.WriteString(fmt.Sprintf("%s Core Maintainers\n\n", maintainersTextTitleStyle))
	for _, m := range mCoreList {
		res.WriteString(fmt.Sprintf("- [%s](%s), %s\n", m.Name, m.Github, m.Company))
	}
	res.WriteString(fmt.Sprintf("\n%s Maintainers\n\n", maintainersTextTitleStyle))
	for _, m := range mList {
		res.WriteString(fmt.Sprintf("- [%s](%s), %s\n", m.Name, m.Github, m.Company))
	}

	return utils.ReplaceTextTags(s, maintainersTextStartTag, maintainersTextEndTag, res.String())
}

var maintainersCmd = &cobra.Command{
	Use:   "maintainers",
	Short: "Generate MAINTAINERS.md for falcosecurity/evolution",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(maintainersReposFilePath) == 0 {
			return fmt.Errorf("must specify a path to repositories.yaml")
		}
		if len(maintainersInFilePath) == 0 {
			return fmt.Errorf("must specify a path to maintainers.yaml")
		}
		if len(maintainersOutFilePath) == 0 {
			return fmt.Errorf("must specify an output markdown file")
		}
		return utils.EditCreateTextFile(maintainersOutFilePath, maintainersTextEditor)
	},
}

func init() {
	maintainersCmd.Flags().StringVarP(&maintainersReposFilePath, "repositories", "r", "", "Path to a repositories.yaml file")
	maintainersCmd.Flags().StringVarP(&maintainersInFilePath, "maintainers", "m", "", "Path to a maintainers.yaml file")
	maintainersCmd.Flags().StringVarP(&maintainersOutFilePath, "output", "o", "", "Path to an output markdown file")
}
