// This is my version of even ended

package main

import (
	"fmt"
)

func main() {
	number := 1112
	numberString := fmt.Sprintf("%d", number)
	// Check if first and last digit is same.
	if numberString[0] == numberString[len(numberString)-1] {
		fmt.Printf("Yes. %v\n", number)
	} else {
		fmt.Printf("No. %v\n", number)
	}

	count := 0
	for a := 1000; a < 10000; a++ {
		for b := a; b < 10000; b++ {
			multiplication := a * b
			multiplicationString := fmt.Sprintf("%d", multiplication)
			// Check if first and last digit is same.
			if multiplicationString[0] == multiplicationString[len(multiplicationString)-1] {
				fmt.Printf("Yes. %v\n", multiplication)
				count++
			}
		}
	}
	fmt.Printf("Count. %v\n", count)

}
