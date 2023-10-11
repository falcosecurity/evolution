// SPDX-License-Identifier: Apache-2.0
/*
Copyright (C) 2023 The Falco Authors.

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
	"os"
	"time"

	"github.com/falcosecurity/evolution/utils/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	latestUpdateTextStartTag = "<!-- LATEST-UPDATE -->"
	latestUpdateTextEndTag   = "<!-- /LATEST-UPDATE -->"
)

var rootCmd = &cobra.Command{
	Use:   "utils",
	Short: "utils - CLI tool for managing falcosecurity/evolution",
	Run: func(c *cobra.Command, args []string) {
		c.Help()
	},
}

func latestUpdateTextEditor(s string) (string, error) {
	if len(s) == 0 {
		s = latestUpdateTextStartTag + latestUpdateTextEndTag
	}

	str := time.Now().UTC().Format(time.RFC3339)
	return utils.ReplaceTextTags(s, latestUpdateTextStartTag, latestUpdateTextEndTag, str)
}

func main() {
	rootCmd.AddCommand(readmeCmd)
	rootCmd.AddCommand(maintainersCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "CLI error: %s\n", err)
		os.Exit(1)
	}
}
