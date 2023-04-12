package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	var counter int
	var bankBalance int = 0
	var balance sync.Mutex

	fmt.Printf("Initial account balance is %d \n", bankBalance)

	incomes := []Income{
		{
			Source: "Main job",
			Amount: 500,
		},
		{
			Source: "Gifts",
			Amount: 100,
		},
		{
			Source: "Part time job",
			Amount: 50,
		},
		{
			Source: "Investments",
			Amount: 100,
		},
	}

	wg.Add(len(incomes))

	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				counter++
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()
				fmt.Printf("On week %d, you earned %d from %s\n", week, income.Amount, income.Source)
			}

		}(i, income)
	}

	wg.Wait()

	fmt.Printf("Final bank balance is %d\n", bankBalance)

	fmt.Println("counter was", counter)
}
