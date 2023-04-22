package main

import "fmt"

type Person interface {
	SayHello()
}

type person struct {
	Name string
	Age  int
}

func (p *person) SayHello() {
	fmt.Printf("My name is %s and I am %d years old", p.Name, p.Age)
}

type tiredPerson struct {
	Name string
	Age  int
}

func (p *tiredPerson) SayHello() {
	fmt.Printf("My name is %s and I am %d years old", p.Name, p.Age)
}

func NewPerson(name string, age int) Person {
	if age > 60 {
		return &tiredPerson{name, age}
	}
	return &person{name, age}
}

func main() {
	// initialize directly
	p := NewPerson("John", 22)
	fmt.Println(p)
	fmt.Printf("%T\n", p)

	p1 := NewPerson("John", 100)
	fmt.Println(p1)
	fmt.Printf("%T\n", p1)

	// use a constructor
	// p2 := NewPerson("Jane", 21)
	// p2.Age = 30 // now this is not allowed
	// fmt.Println(p2)
}
