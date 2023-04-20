package main

import "fmt"

type a struct {
	name string
}

func (val a) print() {
	fmt.Println(val.name)
}

func main() {

	a1 := a{
		name: "eshan",
	}

	a.print(a1)

}
