// My version of for

package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 4; i++ {
		fmt.Println(i)
	}
	fmt.Println()
	for i := 0; i < 4; i++ {
		if i > 2 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println()
	for i := 0; i < 4; i++ {
		if i < 1 {
			continue
		}
		fmt.Println(i)

	}

	i := 0
	fmt.Println()
	for i < 3 {
		fmt.Println(i)
		i++
	}
	fmt.Println()

	j := 0
	for {
		if j > 2 {
			break
		}
		fmt.Println(j)
		j++
	}
}
