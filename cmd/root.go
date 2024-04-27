/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "landmarks",
	Short: "Landmarks API - 地点登録API ",
	Long:  `地点登録APIは、Go言語+OpenAPIの開発サンプルとして作成しました.`,
	Run: func(cmd *cobra.Command, args []string) {
		address := viper.GetString("listen")
		APIMain(address)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.MousetrapHelpText = ""
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "設定ファイル名を指定.デフォルトは、実行ファイル直下のlandmarks-config.yaml")
}

func initConfig() {
	if cfgFile != "" {
		fmt.Fprintln(os.Stderr, "指定された設定ファイルを読み込みます.", cfgFile)
		viper.SetConfigFile(cfgFile)
	} else {
		path, err := os.Executable()
		cobra.CheckErr(err)
		exec := filepath.Dir(path)
		viper.AddConfigPath(exec)
		viper.SetConfigType("yaml")
		viper.SetConfigName("landmarks-config")
		fmt.Fprintln(os.Stderr, "デフォルトの設定ファイルを読み込みます.")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "次の設定ファイルを適用します.", viper.ConfigFileUsed())
	} else {
		fmt.Fprintln(os.Stderr, "設定ファイルが見つかりません.", viper.ConfigFileUsed())
		os.Exit(-1)
	}
}
