package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
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
		go printSomething(val, &wg)
	}

	wg.Wait()

	fmt.Println("main ends here")
}
