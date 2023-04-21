package main

import "fmt"

type a struct {
	name string
}

func (val *a) print() *a {
	fmt.Println(val.name)
	return val
}

func main() {
	a1 := a{
		name: "eshan",
	}

	a1.print().print()
}
