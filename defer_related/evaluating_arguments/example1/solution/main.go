package main

import "fmt"

func main() {
	a := 0
	defer func() {
		fmt.Println(a)
	}()
	a++
}
