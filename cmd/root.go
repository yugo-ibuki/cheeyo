package cmd

import (
	"cheeyo/pkg/cheer"
	"fmt"
	"github.com/spf13/cobra"
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
