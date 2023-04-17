package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if a := os.Args; len(a) != 2 {
		fmt.Println("Give me a number")
	} else if n, err := strconv.Atoi(a[1]); err != nil {
		fmt.Printf("Cannot comvert %q\n", a[1])
	} else {
		fmt.Println("Converted number is ", n)
	}

	// fmt.Println(a, err, n) // all undefined error

}
