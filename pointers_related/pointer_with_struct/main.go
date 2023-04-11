package main

import "fmt"

type eshan struct {
	name string
}

func (e *eshan) updateName(name string) {
	(*e).name = name
}

func main() {
	esh := eshan{
		name: "eshan",
	}

	fmt.Println(esh)

	esh.updateName("eshan1")

	fmt.Println(esh)
}
