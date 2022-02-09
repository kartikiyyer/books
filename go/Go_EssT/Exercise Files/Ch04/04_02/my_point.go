// This is my version of point struct
package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func (p Point) getX() int {
	return p.X
}

func (p *Point) move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}
func main() {
	p := Point{1, 2}
	fmt.Println(p.getX())
	p.move(2, 2)
	fmt.Println(p.X)
}
