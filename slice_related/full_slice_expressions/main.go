package main

import "fmt"

func main() {
	a := []string{"eshan", "eshan1", "eshan2"}

	b := a[:2]

	fmt.Println(len(b))
	fmt.Println(cap(b))

	c := a[:2:2]

	fmt.Println(len(c))
	fmt.Println(cap(c))

	// d := a[:2:1] // error
}
