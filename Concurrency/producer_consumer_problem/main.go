package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

/*
makePizza attempts to make a pizza. We generate a random number from 1-12, and put in two cases where we can't make the pizza
in time. otherwise, we make pizza without issue. To make things interesting, each pizza will take a different length of time to produce
*/
func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1 // delay for atleast 1 second here
		fmt.Printf("Received order %d\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}

		total++

		fmt.Printf("making pizza %d, it will take %d seconds\n", pizzaNumber, delay)

		// delay for a bit
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("We ran out of ingredients for pizza %d", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("The cook quit while making pizza %d", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order %d is ready", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

		return &p
	} else {
		return &PizzaOrder{
			pizzaNumber: pizzaNumber,
		}

	}
}

/*
pizzeria is a goroutine that runs in background and calls makePizza to try to make one order each time it iterates through
the for loop. It executes until it receives something on the quit channel. The quit channel does not receive anything until
the consumer sends it (when the number of orders is greater than or equal to the constant NumberOfPizzas)
*/
func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making
	var i int = 0
	// run forever or until we get quit notification
	// try to make pizza

	for {
		currentPizza := makePizza(i)

		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			// we tried to make a pizza
			case pizzaMaker.data <- *currentPizza:
				{

				}

			case quitChan := <-pizzaMaker.quit:
				{
					close(pizzaMaker.data)
					close(quitChan)
					return
				}
			}
		}
	}
}

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// print out a message
	color.Cyan("Pizzeria is open for business")
	color.Cyan("-----------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in background
	go pizzeria(pizzaJob)

	// create and run consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order %d is out for delivery", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer is really mad")
			}
		} else {
			color.Cyan("Done making pizzas")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("Error closing channel")
			}
		}
	}

	// print out the ending message
	color.Cyan("Done for the day")

	color.Cyan("We made %d pizzas, but failed to make %d, with %d attempts in total", pizzasMade, pizzasFailed, total)
}
