package main

import (
	"fmt"
	"runtime"
	"time"
)

// channel to send square of integers
var c = make(chan int)

// send square of numbers
func square() {
	i := 0

	for {
		i++
		c <- i * i
	}
}

// main goroutine
func main() {

	go square() // start square goroutine

	// get 5 square
	for i := 0; i < 5; i++ {
		fmt.Println("Next square is", <-c)
	}

	// do other job
	time.Sleep(3 * time.Second)

	// print active goroutines
	fmt.Println("Number of active goroutines", runtime.NumGoroutine())
}
