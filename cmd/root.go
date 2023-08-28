package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yugo-ibuki/cheeyo/pkg/cheer"
	"github.com/yugo-ibuki/cheeyo/pkg/words"
	"math/rand"
	"os"
	"time"
)

var rootCmd = &cobra.Command{
	Use:   "this cheers you up",
	Short: "this cheers you up",
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("path").Value.String()

		var w []string
		if len(path) != 0 {
			fmt.Print("path: ", path)
			data, err := os.ReadFile(path)
			if err != nil {
				fmt.Print("error occurs...: ", err)
				os.Exit(1)
			}

			if err := json.Unmarshal([]byte(data), &w); err != nil {
				fmt.Print("error occurs...: ", err)
				os.Exit(1)
			}
		} else {
			w = words.Words
		}

		src := rand.NewSource(time.Now().UnixNano())
		r := rand.New(src)
		// range pickup (0 ~ len(w)-1)
		randomNumber := randIntBetween(r, 0, len(w)-1)
		cheer.Print(w[randomNumber])
	},
}

func init() {
	rootCmd.Flags().StringP("path", "p", "", "words config path")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Print("error occurs...: ", err)
		os.Exit(1)
	}
}

func randIntBetween(r *rand.Rand, min, max int) int {
	return r.Intn(max-min+1) + min
}
