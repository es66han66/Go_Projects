package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

func DoubleHeight(rect *Rectangle) {
	rect.Height = rect.Height * 2
}

func main() {
	rect := Rectangle{
		Width:  10,
		Height: 3,
	}
	fmt.Println("Height before is", rect.Height)

	DoubleHeight(&rect)

	fmt.Println("Height is", rect.Height)
}
