package main

import "fmt"

func main() {
	channel := make(chan string)

	go func() {
		channel <- "eshan"
	}()

	message := <-channel
	fmt.Println(message)
}
