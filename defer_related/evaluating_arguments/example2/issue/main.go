package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) SayHi() {
	fmt.Println("Hi, my name is", p.Name)
}

func main() {
	writer := Person{Name: "Joe"}
	defer writer.SayHi()

	// fix the wrong name
	writer.Name = "Aiden"
}
