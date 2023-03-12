package main

import "fmt"

func Add[T int | float64](a T, b T) T {
	return a + b
}

func main() {
	fmt.Println("Int addition", Add(10, 11))
	fmt.Println("Float addition", Add(1.1, 2.1))
}
