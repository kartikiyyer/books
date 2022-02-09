// This is my version of functions
package main

import (
	"fmt"
)

func divmod(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {
	sum := add(1, 3)
	fmt.Println(sum)
	q, r := divmod(1, 4)
	fmt.Println(q, r)
}

func add(a int, b int) int {
	return a + b
}
