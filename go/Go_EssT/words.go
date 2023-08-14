package main

import (
	"fmt"
	"strings"
)

func main() {
	statement := `
	This is a go lang I am doing a go lang code
	`
	words := strings.Fields(statement)
	fmt.Println(len(words))

	countOfWords := map[string]int{}

	for _, word := range words {
		countOfWords[strings.ToLower(word)]++
	}
	fmt.Println(countOfWords)
	fmt.Printf("%v\n", countOfWords)
}