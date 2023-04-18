package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3}

	a2 := a1[0:2]

	fmt.Println(a1)
	fmt.Println(a2)

	a1[0] = 4

	fmt.Println(a1)
	fmt.Println(a2)
}
