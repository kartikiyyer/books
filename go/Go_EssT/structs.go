package main

import "fmt"

type Center struct {
	X int
	Y int
}

func NewCenter(x int, y int) (*Center, error){
	c := &Center{
		x,
		y,
	}
	return c, nil
}

func (c *Center) Move(dx int, dy int) (int,int) {
	c.X += dx
	c.Y += dy
	return c.X, c.Y
}

type Square struct {
	C *Center
	Length int
}

func NewSquare(x int, y int, length int) (*Square, error){
	c,_ := NewCenter(x, y)
	s := &Square{
		c,
		length,
	}
	return s, nil
}

func (s *Square) Area() int {
	return s.Length * s.Length
}


func main() {
	s,_ := NewSquare(1,2,10)
	fmt.Println(s.C.Move(2,3))
	fmt.Println(s.Area())
}
