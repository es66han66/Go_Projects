package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	rand.Seed(time.Now().UnixNano())

	color.Yellow("The sleeping barber problem")
	color.Yellow("---------------------------")

	/*
		We need a channel for clients in which we can process multiple clients but we cannot process clients more than seating
		capacity hence we will use buffered channel
	*/
	clientChan := make(chan string, seatingCapacity)

	/*
		We need another channel where we can store if cutting is done
	*/
	doneChan := make(chan bool)

	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}

	color.Green("The shop is open for day! \n")

	shop.AddBarber("Frank")
	shop.AddBarber("Susan")

	// we will start barbershop as a go routine
	shopClosing := make(chan bool)
	closed := make(chan bool)
	go func() {
		// we need to run this go routine till the shop closes i.e. shop closing time
		<-time.After(timeOpen)
		/*
			once shop closing time reaches, we will close the shop but because the clients can show up just before the shop closing time
			and can find a seat and wait and barber cannot go until those clients are served hence we first send a message to
			shopClosing and then to closed
		*/
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()

	// add clients
	i := 1
	go func() {
		for {
			// get a random number with average arrival rate
			randomMilliseconds := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				{
					return
				}
			case <-time.After(time.Millisecond * time.Duration(randomMilliseconds)):
				{
					shop.addClient(fmt.Sprintf("Client %d", i))
					i++
				}
			}
		}
	}()

	// we can now just wait for closed channel value, it will only get value when the actual shop is closed completely
	<-closed

}
