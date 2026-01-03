package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	gameId string

	genTsCmd = &cobra.Command{
		Use:   "gen-ts",
		Short: "Generate TypeScript files",
		Long:  `Generate TypeScript files`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("gen-ts")
		},
	}
)

func init() {
	rootCmd.AddCommand(genTsCmd)
	genTsCmd.Flags().StringVar(&gameId, "gameId", "", "Game ID")
	genTsCmd.MarkFlagRequired("gameId")
	viper.BindPFlag("gameId", genTsCmd.Flags().Lookup("gameId"))
}
