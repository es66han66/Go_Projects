package main

import "fmt"

func main() {
	myChannel := make(chan string, 3)
	data := []string{"a", "b", "c"}

	for _, ele := range data {
		select {
		case myChannel <- ele:
		}
	}

	close(myChannel) // if this wasn't here we get, fatal error: all goroutines are asleep - deadlock!

	for result := range myChannel {
		fmt.Println(result)
	}
}
