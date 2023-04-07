package main

import (
	"fmt"
	"os"
)

func main() {
	var data []string

	killswitch := os.Getenv("KILLSWITCH")

	if killswitch == "" {
		fmt.Println("kill switch is off")
		data, err := getData()

		if err != nil {
			panic("ERROR!")
		}

		fmt.Printf("Data was fetched! %d\n", len(data))
	}
	fmt.Println(data)
	for _, item := range data { // here data is not there because the above scope ended within if only
		fmt.Println(item)
	}
}

func getData() ([]string, error) {
	// Simulating getting the data from a datasource - lets say a DB.
	return []string{"there", "are", "no", "strings", "on", "me"}, nil
}
