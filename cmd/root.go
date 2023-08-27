package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yugo-ibuki/cheeyo/pkg/cheer"
	"math/rand"
	"os"
	"time"
)

const (
	dataPath = "./data.json"
)

var rootCmd = &cobra.Command{
	Use:   "this cheers you up",
	Short: "this cheers you up",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := os.ReadFile(dataPath)
		if err != nil {
			fmt.Print("error occurs...: ", err)
			os.Exit(1)
		}

		var words []string
		if err := json.Unmarshal([]byte(data), &words); err != nil {
			fmt.Print("error occurs...: ", err)
			os.Exit(1)
		}

		// シード値の設定（一般的には現在時刻をシードとして使用する）
		src := rand.NewSource(time.Now().UnixNano())
		r := rand.New(src)
		// 10から20の間のランダムな整数を取得
		randomNumber := RandIntBetween(r, 0, len(words))

		cheer.Print(words[randomNumber])
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Print("error occurs...: ", err)
		os.Exit(1)
	}
}

// RandIntBetween は、指定した範囲内でランダムな整数を返します。
func RandIntBetween(r *rand.Rand, min, max int) int {
	return r.Intn(max-min+1) + min
}
