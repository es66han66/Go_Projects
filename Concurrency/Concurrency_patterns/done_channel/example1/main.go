package main

import (
	"fmt"
	"runtime"
	"time"
)

func doWork(done <-chan bool) {
	for {
		select {
		case <-done: // here we are saying that when we get value in done means our go routine has to end
			fmt.Println("end now")
			return
		default: // else our go routine can continue to work
			{
				// fmt.Println("Continue working")
			}
		}
	}
}

func main() {
	done := make(chan bool)

	go doWork(done)

	fmt.Println(runtime.NumGoroutine())

	time.Sleep(5 * time.Second)

	done <- true // or we can close the channel

	// close(done)

	fmt.Println(runtime.NumGoroutine())

}
