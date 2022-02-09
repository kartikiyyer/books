// My version of calculating mean

package main

import (
	"fmt"
)

func main() {

	var a int
	var b int

	a = 1
	b = 2

	fmt.Printf("a=%v, type is %T\n", a, a)
	fmt.Printf("b=%v, type is %T\n", b, b)

	x, y := 5.0, 6.0

	fmt.Printf("x=%v, type is %T\n", x, x)
	fmt.Printf("y=%v, type is %T\n", y, y)

	mean := (x + y) / 2.0

	fmt.Printf("mean=%v, type is %T\n", mean, mean)

}
