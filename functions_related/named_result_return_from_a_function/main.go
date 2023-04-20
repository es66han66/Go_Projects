package main

import "fmt"

func a() (m string) {
	return "eshan"
}

func b() (m string) {
	m = "eshan1"
	return
}

func main() {
	fmt.Println(a())
	fmt.Println(b())
}
