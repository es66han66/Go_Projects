package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	fmt.Println("main started")

	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("1st defer func done")
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("2nd defer func done")
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("3rd defer func done")
	}()

	wg.Wait()

	fmt.Println("main ended")

}
