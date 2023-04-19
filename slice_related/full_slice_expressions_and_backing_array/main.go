package main

import "fmt"

func main() {

	a := []int{1, 2, 3, 4}

	b := a[3:4:4]

	fmt.Println("a is ", a, " with length ", len(a), " and capacity ", cap(a))
	fmt.Println("b is ", b, " with length ", len(b), " and capacity ", cap(b))

	a[3] = 10

	fmt.Println("a is ", a, " with length ", len(a), " and capacity ", cap(a))
	fmt.Println("b is ", b, " with length ", len(b), " and capacity ", cap(b))

	b[0] = 11

	fmt.Println("a is ", a, " with length ", len(a), " and capacity ", cap(a))
	fmt.Println("b is ", b, " with length ", len(b), " and capacity ", cap(b))

	b = append(b, 12) // new backing array created the moment, backing array capacity is exceeded

	fmt.Println("a is ", a, " with length ", len(a), " and capacity ", cap(a))
	fmt.Println("b is ", b, " with length ", len(b), " and capacity ", cap(b))

	a = append(a, 13)

	fmt.Println("a is ", a, " with length ", len(a), " and capacity ", cap(a))
	fmt.Println("b is ", b, " with length ", len(b), " and capacity ", cap(b))
}
