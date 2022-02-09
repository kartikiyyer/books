// My version of error
package main

import (
	"fmt"
	"math"
)

func sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0.0, fmt.Errorf("number is negative")
	}
	return math.Sqrt(num), nil
}

func main() {
	res, err := sqrt(-2.0)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(res)
	}
}
