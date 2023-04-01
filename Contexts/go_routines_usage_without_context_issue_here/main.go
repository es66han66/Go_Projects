package main

import (
	"fmt"
	"runtime"
	"time"
)

func printNumbers() {
	counter := 1

	for {
		time.Sleep(100 * time.Millisecond)
		counter++
	}
}

func main() {
	go printNumbers()

	fmt.Println("Before: active goroutines", runtime.NumGoroutine())
	time.Sleep(time.Second)

	fmt.Println("After: active goroutines", runtime.NumGoroutine())
	fmt.Println("Program exited")
}
