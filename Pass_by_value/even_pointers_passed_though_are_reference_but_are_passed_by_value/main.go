package main

import "fmt"

func main() {
	name := "eshan"

	namePointer := &name

	fmt.Println(namePointer)
	fmt.Println(&namePointer)
	fmt.Println(*namePointer)

	printPointer(namePointer)

}

func printPointer(p *string) {
	fmt.Println(p)
	fmt.Println(&p) // address value is different here
	fmt.Println(*p)
}
