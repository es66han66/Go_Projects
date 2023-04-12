package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

func main() {
	sq := square{
		sideLength: 10,
	}

	tr := triangle{
		base:   10,
		height: 10,
	}

	printArea(sq)

	printArea(tr)

}
