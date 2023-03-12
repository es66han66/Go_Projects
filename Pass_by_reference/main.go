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

	DoubleHeight(&rect)

	fmt.Println("Width still is", rect.Width)
	fmt.Println("Height still is", rect.Height)
}
