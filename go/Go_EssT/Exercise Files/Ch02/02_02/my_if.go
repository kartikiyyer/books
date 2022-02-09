// My version of if

package main

import (
	"fmt"
)

func main() {
	x := 5

	if x > 2 {
		fmt.Println("X is big.")
	}

	if x > 10 {
		fmt.Println("X is very big.")
	} else {
		fmt.Println("X is not that big.")
	}

	if x > 2 && x < 10 {
		fmt.Println("X is just right.")
	}

	if x < 10 || x > 20 {
		fmt.Println("X is out of range.")
	}

	a := 11.0
	b := 20.0

	if frac := a / b; frac > 0.5 {
		fmt.Printf("a is bigger that half of b. %v", frac)
	}
}
