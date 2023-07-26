/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package set

import (
	"os"

	"github.com/hossted/k8/cmd/set/monitoring"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var SetCmd = &cobra.Command{
	Use:     "set",
	Short:   "[s] Change application settings",
	Long:    `[s] Change application settings`,
	Aliases: []string{"s"},
	Example: `
  hossted set monitoring true
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		return nil
	},
}

func init() {
	SetCmd.AddCommand(monitoring.SetMonitoring)
}
