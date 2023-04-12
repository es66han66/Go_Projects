package main

import "fmt"

type eshan struct {
	name string
}

func (eshan) returnWithStaticValue() string {
	return "me"
}

/*
	In this case as we don't use any value from the type eshan hence we didn't provide a name to receiver
*/

func main() {
	var e eshan = eshan{
		name: "eshan",
	}

	fmt.Println(e.returnWithStaticValue())
}
