package main

import "fmt"

func labeledBreakUsage(a, b []string) {

	// without label
	// for i, val := range a {
	// 	for _, val1 := range b {
	// 		if val1 == val {
	// 			fmt.Println("Found ", val1, " from ", b, " in ", a, " at ", i+1)
	// 			break
	// 		}
	// 	}
	// }

	// with label

label:
	for i, val := range a {
		for _, val1 := range b {
			if val1 == val {
				fmt.Println("Found ", val1, " from ", b, " in ", a, " at ", i+1)
				break label
			}
		}
	}
}

func labelsAndVariablesSameNameCase(a, b []string) {
	var label int = 10
	fmt.Println("label at first is", label)
label:
	for i, val := range a {
		for _, val1 := range b {
			if val1 == val {
				fmt.Println("label in between is", label)
				fmt.Println("Found ", val1, " from ", b, " in ", a, " at ", i+1)
				break label
			}
		}
	}
	fmt.Println("label at last is", label)

}

func main() {
	a := []string{"eshan", "eshan1", "eshan", "eshan1"}
	b := []string{"eshan1", "eshan"}

	labeledBreakUsage(a, b)

	labelsAndVariablesSameNameCase(a, b)

}
