package main

import "fmt"

type ShapeCalculator interface {
	Area() int
	Perimeter() int
}

type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}

// uncomment the following line to guarantee that Implementation implements all methods of SomeInterface
// var _ ShapeCalculator = (*Rectangle)(nil) // ← this is the line

func main() {
	rect := Rectangle{
		Width:  10,
		Height: 3,
	}
	fmt.Println(rect.Area())
}
