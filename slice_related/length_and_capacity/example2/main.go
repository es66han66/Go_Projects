package main

import "fmt"

func main() {
	original := [5]int{1, 2, 3}

	fmt.Println(len(original))
	fmt.Println(cap(original))

	s := original[:]

	fmt.Println(len(s))
	fmt.Println(cap(s))

	s = append(s, 5)

	fmt.Println(len(s))
	fmt.Println(cap(s))
}
