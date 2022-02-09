// This is my version of creating new objects
package main

import (
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	Gender string
}

func NewPerson(name string, age int, gender string) (*Person, error) {
	person := &Person{
		name,
		age,
		gender,
	}
	return person, nil
}

func main() {
	p, _ := NewPerson("Kar", 12, "Male")
	fmt.Println(p)
}
