package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}

	fmt.Println(len(a))
	fmt.Println(cap(a))

	a1 := a[0:0]

	fmt.Println(len(a1))
	fmt.Println(cap(a1))

	a2 := a[1:1]

	fmt.Println(len(a2))
	fmt.Println(cap(a2))

}
