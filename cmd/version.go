package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "This command return version number.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version is 1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
