package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *BarberShop) AddBarber(barber string) {
	// we will increment NumberOfBarbers by 1 whenever we add a barber
	shop.NumberOfBarbers++

	/*
		We will firstly launch a go routine where we will first mark the barber as not asleep because once barber enters the shop,
		he is awake. After that barber will go to waiting room and look for whether client is available or not.
	*/
	go func() {
		isSleeping := false

		color.Yellow("%s goes to the waiting room to check for clients.", barber)

		for {
			// if there are no clients, barber goes to sleep
			if len(shop.ClientsChan) == 0 {
				color.Yellow("There is nothing to do, so %s takes a nap", barber)
				isSleeping = true
			}

			// if above is false which means there are client(s), hence we will keep listening to clients channel
			/*
				whenever we receive a client, we first need to check whether shop is opened or not, because we cannot accept a barber
				if client has come after the shop is closed. To check this, we can use Open field in barberShop struct but in that case
				there can be a possible race condition when there will be multiple barbers so we use below approach.
			*/
			client, shopOpen := <-shop.ClientsChan

			if shopOpen {

				/*
					when customer arrives and the barber is sleeping, he needs to wake up the barber
				*/
				if isSleeping {
					color.Yellow("%s wakes %s up", client, barber)
					isSleeping = false
				}

				// cut hair
				shop.cutHair(barber, client)
			} else {
				// shop is closed so send the barber home and close this goroutine
				shop.sendBarberHome(barber)
				return
			}

		}

	}()
}

func (shop *BarberShop) cutHair(barber, client string) {
	color.Green("%s is cutting %s hair.", barber, client)
	// here we will sleep (cut hair) for hair cut time
	time.Sleep(shop.HairCutDuration)
	color.Green("%s is finished cutting %s's hair", barber, client)
}

func (shop *BarberShop) sendBarberHome(barber string) {
	color.Cyan("%s is going home", barber)
	shop.BarbersDoneChan <- true
}

func (shop *BarberShop) closeShopForDay() {
	color.Cyan("Closing shop for the day")
	close(shop.ClientsChan)
	shop.Open = false

	// we need to wait here for all barbers to finish to completely shut shop and send all barbers home
	for a := 1; a <= shop.NumberOfBarbers; a++ {
		// below will block until all barbers send done to their done channels
		<-shop.BarbersDoneChan
	}

	close(shop.BarbersDoneChan)

	color.Green("The barber shop is closed for day and everyone has gone home")
}

func (shop *BarberShop) addClient(client string) {
	color.Green("%s client arrives", client)

	if shop.Open {
		/*
			If shop is opened then we need to check if we have the capacity in clients Channel to accomodate the client
		*/
		select {
		case shop.ClientsChan <- client:
			{
				color.Yellow("%s takes a seat in the waiting room", client)
			}
		default:
			{
				// flow coming here means ClientsChan cannot take any value as it's already full
				color.Red("The waiting room is full, so %s needs to leave", client)
			}
		}
	} else {
		color.Red("Shop is already closed so %s leaves.", client)
	}
}
