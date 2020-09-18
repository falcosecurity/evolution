package lib

import "github.com/spf13/cobra"

var rootCmd = cobra.Command{
	Use: "kilt-test",
	Short: "kilt-test allows you to test kilt-lib",
}

func Execute() error {
	return rootCmd.Execute()
}