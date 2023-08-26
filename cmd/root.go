package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yugo-ibuki/cheeyo/pkg/cheer"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "this cheers you up",
	Short: "this cheers you up",
	Run: func(cmd *cobra.Command, args []string) {
		cheer.Print("頑張れ！！！")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Print("error occurs...: ", err)
		os.Exit(1)
	}
}
