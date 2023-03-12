package main

import "fmt"

func check(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("int type, value is", a.(int))
	case bool:
		fmt.Println("bool type, value is", a.(bool))
	}
}

func main() {
	var a = 10
	check(a)

	var b = "eshan"
	check(b)
}
