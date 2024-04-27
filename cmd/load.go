/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"landmarks/pkg/database"
	"log/slog"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// loadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		if inputFile, err := cmd.Flags().GetString("file"); err != nil {
			slog.Error(err.Error())
		} else {
			dbFile := viper.GetString("database")
			if err := database.LoadDatabase(dbFile, inputFile); err != nil {
				slog.Error(err.Error())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(loadCmd)
	loadCmd.Flags().StringP("file", "f", "", "データJSONファイル")
	loadCmd.MarkFlagRequired("file")
}
