package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var SemVer = "0.0.0-git"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of the command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(SemVer)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
