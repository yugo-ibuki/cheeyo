package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yugo-ibuki/cheeyo/pkg"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "this cheers you up",
	Short: "this cheers you up",
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("path").Value.String()
		merge := cmd.Flag("merge").Value.String()

		w := pkg.NewWord()
		if len(path) != 0 {
			data, err := pkg.ReadConfig(path)
			if err != nil {
				fmt.Print("error occurs...: ", err)
				os.Exit(1)
			}

			// merge words config and existing words
			if merge == "true" {
				w.MergeConfig(data)
			}
		}

		pkg.Print(w.GetRandomOne())
	},
}

func init() {
	rootCmd.Flags().StringP("path", "p", "", "words config path")
	rootCmd.Flags().BoolP("merge", "m", false, "merge words config")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Print("error occurs...: ", err)
		os.Exit(1)
	}
}
