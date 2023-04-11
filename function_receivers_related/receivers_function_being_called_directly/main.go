package main

import "fmt"

type eshan string

func (e eshan) print1() {
	fmt.Println("eshan")
}

func main() {
	var esh eshan = "eshan"
	esh.print1()

	print1() // we get print1 undefined, not allowed
}
