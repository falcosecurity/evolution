package lib

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	kiltlib "lib"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build an configuration object from HOCON",
	Long:  "imports a kilt patch file and outputs parameters based on input",
	Run:   buildKilt,
}

var (
	containerInput string
	kiltPatch string
)

func init() {
	buildCmd.Flags().StringVarP(&containerInput, "container-input", "i", "" , "a json that represents a TargetInfo")
	buildCmd.MarkFlagRequired("container-input")
	buildCmd.Flags().StringVarP(&kiltPatch, "kilt", "k", "", "kilt definition to process")
	buildCmd.MarkFlagRequired("kilt")

	rootCmd.AddCommand(buildCmd)
}

func buildKilt(cmd *cobra.Command, args []string)  {
	containerInputData, err := ioutil.ReadFile(containerInput)
	if err != nil {
		panic(err)
	}

	input := new(kiltlib.TargetInfo)
	err = json.Unmarshal(containerInputData, input)
	if err != nil {
		panic(err)
	}

	k, err := kiltlib.KiltFromFile(input, kiltPatch)

	if err != nil {
		panic(err)
	}

	b, err := k.Build()

	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", b)
	if k.HasRuntime() {
		r, err := k.Runtime()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", r)
	}

}

