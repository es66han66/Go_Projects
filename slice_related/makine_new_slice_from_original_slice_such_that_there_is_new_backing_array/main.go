package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}

	// b := append([]int(nil), a[:]...) or below
	b := append([]int{}, a[:]...)

	fmt.Println(a)
	fmt.Println(b)

	b[0] = 5

	fmt.Println(a)
	fmt.Println(b)

}
