package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	orgFile  string
	destFile string

	rootCmd = &cobra.Command{
		Use:   "jsonToModel",
		Short: "json 模型自动转成其它语言的数据模型工具",
		Long:  `这是一个将json 模型自动转成其它语言的数据模型工具：如 .go ; .dart`,
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.scaffold.yaml)")
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.AddConfigPath(home)
	viper.SetConfigName(".scaffold")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
