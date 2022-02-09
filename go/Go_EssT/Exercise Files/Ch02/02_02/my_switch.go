// My version of switch
package main

import (
	"fmt"
)

func main() {
	x := 4
	switch x {
	case 1:
		fmt.Println("X is 1")
	case 2:
		fmt.Println("X is 2")
	default:
		fmt.Println("X is more")
	}

	switch {
	case x > 100:
		fmt.Println("X is greater than 100.")
	default:
		fmt.Println("X is less than 100.")
	}

}
