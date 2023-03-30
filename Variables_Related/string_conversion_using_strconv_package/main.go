package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 42
	var j string = strconv.Itoa(i)
	fmt.Printf("%v %T \n", i, i)
	fmt.Printf("%v %T", j, j)
}
