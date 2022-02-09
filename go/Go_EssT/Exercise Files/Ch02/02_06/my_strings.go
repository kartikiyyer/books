// This is my version of strings
package main

import (
	"fmt"
)

func main() {
	statement := "Hi welcome to Go!"

	fmt.Println(statement)
	fmt.Println("Statement: " + statement)
	fmt.Println(len(statement))
	fmt.Printf("%v : %T\n", statement[0], statement[0])
	// Returns bytes of the character accessed.
	fmt.Println(statement[0])

	// Returns a string
	fmt.Println(statement[0:1])
	fmt.Println(statement[4:11])
	fmt.Println(statement[4:])
	fmt.Println(statement[:11])

	multiline := `
	This is a multiline 
	comment.
	This is another line.
	`
	fmt.Println(multiline)
}
