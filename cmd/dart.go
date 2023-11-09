package cmd

import (
	"fmt"
	"github.com/any-call/gobase/util/myos"
	"github.com/jinguihua/jsonToModel/todart"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(dartCmd)

	dartCmd.PersistentFlags().StringVar(&orgFile, "orgfile", "", "请输入源文件地址")
	dartCmd.PersistentFlags().StringVar(&destFile, "destfile", "", "请输入目前文件地址")
}

var dartCmd = &cobra.Command{
	Use:   "dart",
	Short: "转成dart 语言模型",
	Long:  `读取 json 数据,将结构自动转成 dart 语法模型`,
	Run: func(cmd *cobra.Command, args []string) {
		if orgFile == "" {
			fmt.Fprintln(os.Stderr, "empty origin file path")
			os.Exit(1)
		}

		if destFile == "" {
			fmt.Fprintln(os.Stderr, "empty destination file path")
			os.Exit(1)
		}

		//检测文件路径格式
		if b := myos.IsExistFile(orgFile); !b {
			fmt.Fprintln(os.Stderr, "origin file isn't exist at:", orgFile)
			os.Exit(1)
		}

		jsonBytes, err := os.ReadFile(orgFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "invalid org file :", orgFile)
			os.Exit(1)
		}

		if err := todart.GenModel(destFile, string(jsonBytes)); err != nil {
			fmt.Fprintln(os.Stderr, "generate dart file fail:", err.Error())
			os.Exit(1)
		} else {
			fmt.Fprintln(os.Stdout, "generate dart file successful :", destFile)
		}
		os.Exit(0)
	},
}
