// This is my version of square
package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

type Square struct {
	Point  Point
	Length int
}

func NewSquare(x int, y int, length int) (*Square, error) {
	point := Point{x, y}
	square := &Square{
		Point:  point,
		Length: length,
	}
	return square, nil
}

func (s *Square) Move(dx int, dy int) {
	s.Point.Move(dx, dy)
}
func (s *Square) Area() int {
	return s.Length * s.Length
}

func main() {
	square, _ := NewSquare(1, 2, 4)
	fmt.Println(square)
	square.Move(2, 3)
	fmt.Println(square)
	fmt.Println(square.Area())
}
