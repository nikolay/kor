package kor

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yonahd/kor/pkg/kor"
)

var stsCmd = &cobra.Command{
	Use:     "statefulsets",
	Aliases: []string{"sts"},
	Short:   "Gets unused statefulSets",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if outputFormat == "json" || outputFormat == "yaml" {
			if response, err := kor.GetUnusedStatefulSetsStructured(includeExcludeLists, kubeconfig, outputFormat); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(response)
			}
		} else {
			kor.GetUnusedStatefulSets(includeExcludeLists, kubeconfig)
		}

	},
}

func init() {
	rootCmd.AddCommand(stsCmd)
}
