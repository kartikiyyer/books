// My version of slices
package main

import (
	"fmt"
)

func main() {
	strings := []string{"Adam", "Eve", "Jack"}
	fmt.Println(strings)
	fmt.Println(strings[0])
	fmt.Println(strings[0:2])
	fmt.Println(len(strings))

	var againStrings [5]string
	againStrings[0] = "Christian"
	againStrings[4] = "New"

	fmt.Println(againStrings)

	for i := range strings {
		fmt.Println(i)
		fmt.Println(strings[i])
	}

	for i, name := range strings {
		fmt.Println(i)
		fmt.Println(name)

	}
	for _, name := range strings {
		fmt.Println(name)

	}

	strings1 := append(strings, "Jill")
	fmt.Println(strings1)
	fmt.Println(strings)
}
