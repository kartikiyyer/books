package main

import (
	"fmt"
)

func main() {
	numbers := []int{16, 8, 42, 4, 23, 15}
	var max int
	for _, number := range numbers {
		if max < number {
			max = number
		}
	}
	fmt.Println(max)

	num := 2.0
	fmt.Printf("%.2f", num)
}