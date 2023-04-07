package main

import "fmt"

func main() {
	channel := make(chan string)
	channel1 := make(chan string)

	go func() {
		channel <- "eshan"
	}()

	go func() {
		channel1 <- "eshan1"
	}()

	select {
	case channelMessage := <-channel:
		{
			fmt.Println(channelMessage)
		}

	case channel1Message := <-channel1:
		{
			fmt.Println(channel1Message)
		}
	}
	fmt.Println("main ends here")
}
