package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func main() {
	msg = "hello world"

	wg.Add(2)

	go updateMessage("eshan")

	go updateMessage("eshan1")

	wg.Wait()

	fmt.Println(msg)
}
