package main

import (
	"fmt"
)

func main() {
	var data []string

	data, err := getData() // here the scope is main i.e. why we get the value in data uptil whole main function
	if err != nil {
		panic("ERROR!")
	}

	for _, item := range data {
		fmt.Println(item)
	}
}

func getData() ([]string, error) {
	// Simulating getting the data from a datasource - lets say a DB.
	return []string{"there", "are", "no", "strings", "on", "me"}, nil
}
