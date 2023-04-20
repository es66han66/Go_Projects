package main

import "fmt"

type int1 interface {
	print()
}

type cust1 struct {
	name string
}

func (val cust1) print() {
	fmt.Println(val.name)
}

func (val cust1) printExtra() {
	fmt.Println("Extra ", val.name)
}

type cust2 struct {
	name string
}

func (val cust2) print() {
	fmt.Println(val.name)
}

func receiveInt1SatisfyingInterfaceWithPrintExtraMethodAttached(val int1) {
	type checkPrintExtra interface {
		printExtra()
	}
	if it, ok := val.(checkPrintExtra); ok {
		it.printExtra()
	} else {
		fmt.Println("Doesn't have printExta attached")
	}
}

func receiveInt1SatisfyingInterface(val int1) {
	val.print()
}

func main() {

	var_cust1 := cust1{
		name: "eshan",
	}

	var_cust2 := cust2{
		name: "eshan1",
	}

	receiveInt1SatisfyingInterface(var_cust1)

	receiveInt1SatisfyingInterface(var_cust2)

	receiveInt1SatisfyingInterfaceWithPrintExtraMethodAttached(var_cust1)

	receiveInt1SatisfyingInterfaceWithPrintExtraMethodAttached(var_cust2)

}
