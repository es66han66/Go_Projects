package main

import "fmt"

func returnValue() string {
	return "hello"
}

func main() {
	const variable string = returnValue() // error- returnValue() (value of type string) is not constant
	// const variable string = "hello"       // no error
	fmt.Println(variable)
}
