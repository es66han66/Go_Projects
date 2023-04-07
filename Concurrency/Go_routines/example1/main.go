package main

import (
	"fmt"
	"time"
)

func someFunc() {
	fmt.Println("1")
}

func main() {
	go someFunc()
	go someFunc()

	time.Sleep(5 * time.Second)

	fmt.Println("main ended")
}
