package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	msg = s
	m.Unlock()
}

func main() {
	msg = "hello world"
	var mutex sync.Mutex

	wg.Add(2)

	go updateMessage("eshan", &mutex)

	go updateMessage("eshan1", &mutex)

	wg.Wait()

	fmt.Println(msg)
}
