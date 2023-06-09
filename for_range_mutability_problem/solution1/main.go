package main

import (
	"fmt"
	"sync"
)

type A struct {
	id int
}

func main() {
	channel := make(chan A, 5)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for a := range channel {
			wg.Add(1)
			go func(item A) {
				defer wg.Done()
				fmt.Println(item.id)
			}(a)
		}

	}()

	for i := 0; i < 10; i++ {
		channel <- A{id: i}
	}
	close(channel)

	wg.Wait()
}
