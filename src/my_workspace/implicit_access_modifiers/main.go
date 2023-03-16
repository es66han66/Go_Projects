package main

import (
	"fmt"

	hello "my_workspace/implicit_Access_Modifiers/mylib"
)

func main() {
	fmt.Println(hello.TryExporting)
	// fmt.Println(hello.tryExporting) // will give error
}
