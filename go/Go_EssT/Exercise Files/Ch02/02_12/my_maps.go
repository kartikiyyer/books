// My version of maps
package main

import (
	"fmt"
)

func main() {
	stocks := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	fmt.Println(stocks)

	for key := range stocks {
		fmt.Println(key)
	}
	for key, value := range stocks {
		fmt.Println(key, value)
	}

	fmt.Println(stocks["D"])

	value, ok := stocks["D"]

	fmt.Println(value, ok)

	stocks["D"] = 4

	fmt.Println(stocks)

	delete(stocks, "D")
	fmt.Println(stocks)
}
