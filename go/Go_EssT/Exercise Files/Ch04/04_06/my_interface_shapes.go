// This is my version of interface shapes
package main

import (
	"fmt"
	"math"
)

type Square struct {
	Length float64
}

func (s *Square) Area() float64 {
	return s.Length * s.Length
}

type Circle struct {
	Radius float64
}

func (s *Circle) Area() float64 {
	return math.Pi * s.Radius * s.Radius
}

type Shape interface {
	Area() float64
}

func sumArea(shapes []Shape) float64 {
	var sum float64
	for _, shape := range shapes {
		sum += shape.Area()
	}
	return sum
}

func main() {
	square := &Square{2.0}
	circle := &Circle{3.0}
	fmt.Println(sumArea([]Shape{square, circle}))
}
