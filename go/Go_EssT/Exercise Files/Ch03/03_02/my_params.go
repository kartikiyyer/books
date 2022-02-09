// My version of parameter passing
package main

import (
	"fmt"
)

func doubleAt(numbers []int, i int) {
	numbers[i] *= 2
}

func multiply(number int) {
	number *= 2
}

func multiplyPtr(number *int) {
	*number *= 2
}

func main() {
	numbers := []int{1, 2, 4, 5}
	doubleAt(numbers, 1)
	fmt.Println(numbers)

	number := 2
	multiply(number)
	fmt.Println(number)

	multiplyPtr(&number)
	fmt.Println(number)
}
