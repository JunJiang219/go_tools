package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of mycli",
	Long:  `All software has versions. This is mycli's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mycli version 1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
