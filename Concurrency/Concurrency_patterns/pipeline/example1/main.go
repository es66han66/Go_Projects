package main

import "fmt"

func putNumsToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, val := range nums {
			out <- val
		}
		close(out)
	}()
	return out
}

func squareAllValuesOnChannel(input <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for val := range input {
			out <- val * val
		}
		close(out)
	}()
	return out
}

func main() {
	numSlice := []int{1, 2, 3, 4}

	inputChannel := putNumsToChannel(numSlice)

	squaredChannel := squareAllValuesOnChannel(inputChannel)

	for val := range squaredChannel {
		fmt.Println(val)
	}
}
