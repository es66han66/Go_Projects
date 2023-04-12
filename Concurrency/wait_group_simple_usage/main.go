package main

import (
	"fmt"
	"sync"
)

func printSomething(s string) {
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup
	words := []string{
		"eshan",
		"hello",
		"address",
	}

	for _, val := range words {
		wg.Add(1)
		go func(val string) {
			defer wg.Done()
			printSomething(val)
		}(val)
	}

	wg.Wait()

	fmt.Println("main ends here")
}
