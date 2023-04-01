package main

import (
	"fmt"
	"runtime"
	"time"
)

var signal = make(chan bool)

func printNumbers() {
	counter := 1

	for {
		select {
		case <-signal:
			return
		default:
			time.Sleep(100 * time.Millisecond)
			counter++
		}
	}
}

func main() {
	go printNumbers()

	fmt.Println("Before: active goroutines", runtime.NumGoroutine())
	time.Sleep(time.Second)

	signal <- true

	fmt.Println("After: active goroutines", runtime.NumGoroutine())
	fmt.Println("Program exited")
}
