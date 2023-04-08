package main

import "fmt"

func main() {
	exampleSlice := []int{0, 1, 2, 3, 4}
	newSlice := exampleSlice[1:2]
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
}
