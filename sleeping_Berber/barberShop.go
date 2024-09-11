package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberSop struct {
	NumOfSeatCapacity int
	NumOfBarbers      int
	cutttingDuration  time.Duration
	BarberDoneChan    chan bool
	clientChan        chan string
	Open              bool
}

func (shop *BarberSop) addBarber(barber string) {
	shop.NumOfBarbers++
	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room  check for clients", barber)
		for {
			//if there are no clients the barber going to sleep
			if len(shop.clientChan) == 0 {
				color.Yellow("There is nothing to do. So %s is sleeping in the chair", barber)
				isSleeping = true

			}
			client, shopOpen := <-shop.clientChan
			//second return parameter we get  when reading a channel
			//wheater the value recived from the channel was actually sent to that channel
			//true /false....which means the channel is closed and empty

			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up.", client, barber)
					isSleeping = false
				}
				//cut hair..........
				shop.cutHair(barber, client)

			} else {

				// shop is closed, so send the barber home and close this goroutine
				shop.sendBerberHome(barber)
				return
			}

		}

	}()
}
func (shop *BarberSop) cutHair(barber, client string) {
	color.Green("% is cutting %s's hair.", barber, client)
	time.Sleep(shop.cutttingDuration)
	color.Green("%s is finished cutting %s's hair.", barber, client)
}
func (shop *BarberSop) sendBerberHome(barber string) {
	color.Cyan("%s is going home.", barber)

	shop.BarberDoneChan <- true

}
func (shop *BarberSop) closeBarberForDay() {
	color.Cyan("closing shop for day...")
	close(shop.clientChan)
	shop.Open = false
	for i := 0; i < shop.NumOfBarbers; i++ {
		<-shop.BarberDoneChan
		//this will wait until every single barber is done

	}
	close(shop.BarberDoneChan)

}
