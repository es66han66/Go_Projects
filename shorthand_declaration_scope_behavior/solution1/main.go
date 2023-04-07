package main

import (
	"fmt"
	"os"
)

func main() {
	var data []string
	var err error // Declaring err to make sure we can use = instead of :=

	killswitch := os.Getenv("KILLSWITCH")

	if killswitch == "" {
		fmt.Println("kill switch is off")
		data, err = getData()

		if err != nil {
			panic("ERROR!")
		}

		fmt.Printf("Data was fetched! %d\n", len(data))
	}

	for _, item := range data {
		fmt.Println(item)
	}
}

func getData() ([]string, error) {
	// Simulating getting the data from a datasource - lets say a DB.
	return []string{"there", "are", "no", "strings", "on", "me"}, nil
}
