package main

import "fmt"

type myStruct struct {
	name string
}

func updateSlice(s []int) {
	s[0] = 10
}

func updateStruct(s myStruct) {
	s.name = "eshan1"
}

func main() {
	mySlice := []int{1, 2, 3, 4}
	fmt.Println("slice before", mySlice)
	updateSlice(mySlice)
	fmt.Println("slice after", mySlice)

	structVar := myStruct{
		name: "eshan",
	}
	fmt.Println("struct before", structVar)

	updateStruct(structVar)
	fmt.Println("struct after", structVar)

}

/*
	Slices when passed result in automatic pass by reference without pointers, while struct are not
*/
