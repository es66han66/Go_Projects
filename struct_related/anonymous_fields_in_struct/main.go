package main

import "fmt"

func main() {
	type a struct {
		name string
	}

	type b struct {
		address string
		a
	}

	a1 := b{
		address: "lucknow",
		a: a{
			name: "eshan",
		},
	}

	fmt.Println(a1.a.name)
	fmt.Println(a1.name)

}
