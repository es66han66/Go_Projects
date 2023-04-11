package main

import "fmt"

type eshan string

func (e eshan) printMe() {
	fmt.Println(e)
}

func main() {
	var esh eshan = "eshan"
	esh.printMe()
}
