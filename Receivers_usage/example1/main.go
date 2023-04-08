package main

import "fmt"

type A struct {
	name string
}

func (a A) GetName() string {
	return a.name
}

func main() {
	typeA := A{
		name: "eshan",
	}

	fmt.Println(typeA.GetName())
}
