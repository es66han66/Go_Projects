package main

import "fmt"

type eshan []string

func (e eshan) print() {
	e = append(e, "hello")
}

func main() {
	var esh eshan = []string{"eshan"}
	fmt.Println(esh)
	esh.print()
	fmt.Println(esh)
}
