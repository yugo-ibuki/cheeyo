package pkg

import (
	"fmt"
	"time"
)

func Print(dialog string) {
	var upperDisplay = []rune("＿人人人＿")
	var downDisplay = []rune("￣Y^Y^Y^Y￣")
	display := ""
	for i, v := range dialog {
		// skip first time
		if i != 0 {
			clear()
		}

		var upperTemp = string(upperDisplay)
		var downTemp = string(downDisplay)

		// if even, replace the last character or if over 15, replace the last character.
		// because appending upper and down is faster than replacing the cheering word.
		if i%2 == 0 || i > 15 {
			upperTemp = addLength(upperDisplay, "人")
			downTemp = addLength(downDisplay, "^Y")
		}

		printFormatted(upperTemp, downTemp, display)

		upperDisplay = []rune(upperTemp)
		downDisplay = []rune(downTemp)
		display += string(v)
		time.Sleep(100 * time.Millisecond)
	}
}

// clear - clear the dialog
func clear() {
	fmt.Print("\033[1A\033[K") // first line
	fmt.Print("\033[1A\033[K") // second line
	fmt.Print("\033[1A\033[K") // third line
}

func printFormatted(upperTemp, downTemp, display string) {
	fmt.Println(upperTemp)
	fmt.Printf("＞ %s  ＜", display)
	fmt.Println()
	fmt.Println(downTemp)
}

// addLength - add length to the last character
func addLength(runeWord []rune, word string) string {
	return string(runeWord[:len(runeWord)-1]) + word + string(runeWord[len(runeWord)-1:])
}
