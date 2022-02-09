// My version of couting the words.

package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needles and pins
	Needles and pins
	Swe me a sail
	To catch me the wind.
	`

	words := strings.Fields(text)

	count := map[string]int{}

	for _, word := range words {
		fmt.Println(strings.ToLower(word))
		if _, ok := count[strings.ToLower(word)]; ok {
			count[strings.ToLower(word)]++
		} else {
			count[strings.ToLower(word)] = 1
		}
	}
	fmt.Println(count)
}
