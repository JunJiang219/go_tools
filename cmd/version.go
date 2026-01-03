package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of mycli",
	Long:  `All software has versions. This is mycli's`,
	Run: func(cmd *cobra.Command, args []string) {
		version := viper.GetString("version")
		fmt.Printf("mycli version %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
