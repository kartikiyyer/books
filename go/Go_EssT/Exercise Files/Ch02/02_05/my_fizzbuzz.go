// My version of fizzbuzz

package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("fizzbuzz %v\n", i)
		} else if i%3 == 0 {
			fmt.Printf("fizz %v\n", i)
		} else if i%5 == 0 {
			fmt.Printf("buzz %v\n", i)
		} else {
			fmt.Println(i)
		}
	}
}
