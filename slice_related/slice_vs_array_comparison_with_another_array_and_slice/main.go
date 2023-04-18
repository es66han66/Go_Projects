package main

import "fmt"

func main() {
	a1 := []int{1, 2}
	a2 := []int{1, 2}

	fmt.Println(a1 == a2) // error

	a3 := [2]int{1, 2}
	a4 := [2]int{3, 4}

	fmt.Println(a3 == a4) // no error allowed

}
