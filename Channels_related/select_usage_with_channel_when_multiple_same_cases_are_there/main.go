package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	for {
		time.Sleep(5 * time.Second)

		ch <- "This is from server1"
	}
}

func server2(ch chan string) {
	for {
		time.Sleep(3 * time.Second)

		ch <- "This is from server2"
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go server1(ch1)
	go server2(ch2)
	for {
		select {
		case s := <-ch1:
			{
				fmt.Println("Message from server1 case 1 is ", s)
			}
		case s1 := <-ch1:
			{
				fmt.Println("Message from server1 case 2 is ", s1)
			}
		case s3 := <-ch2:
			{
				fmt.Println("Message from server1 case 3 is ", s3)
			}
		case s4 := <-ch2:
			{
				fmt.Println("Message from server1 case 4 is ", s4)
			}
		default:
			{
				fmt.Println("End day")
				break
			}
		}
	}
}

/*
	In the situation where there are multiple cases with same condition then any random is picked
*/
