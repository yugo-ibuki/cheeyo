package pkg

import (
	"encoding/json"
	"github.com/yugo-ibuki/cheeyo/assets"
	"math/rand"
	"os"
	"time"
)

type Word struct {
	Words []string
}

func NewWord() *Word {
	return &Word{
		Words: assets.Words,
	}
}

// GetRandomOne - get random one word
func (w *Word) GetRandomOne() string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	// range pickup (0 ~ len(w)-1)
	randInt := randIntBetween(r, 0, len(w.Words)-1)
	return w.Words[randInt]
}

// MergeConfig - merge config words and existing words
func (w *Word) MergeConfig(configWords []string) {
	w.Words = append(w.Words, configWords...)
}

// randIntBetween - get random int between min and max
func randIntBetween(r *rand.Rand, min, max int) int {
	return r.Intn(max-min+1) + min
}

// ReadConfig - read config json file
func ReadConfig(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var configWords []string
	if err := json.Unmarshal(data, &configWords); err != nil {
		return nil, err
	}

	return configWords, nil
}
