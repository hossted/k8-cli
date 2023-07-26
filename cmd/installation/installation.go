/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package installation

import (
	"os"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var InstallationCmd = &cobra.Command{
	Use:     "set",
	Short:   "[s] Change application settings",
	Long:    `[s] Change application settings`,
	Aliases: []string{"s"},
	Example: `
  hossted init
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

}
