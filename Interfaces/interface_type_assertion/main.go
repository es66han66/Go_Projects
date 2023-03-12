package main

import "fmt"

func check(a interface{}) {
	value, ok := a.(string)
	fmt.Println(ok, " ", value)
}

func main() {
	var a interface{} = 10
	check(a)

	var b interface{} = "Eshan"
	check(b)
}
