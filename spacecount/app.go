package main

import . "github.com/alexellis/blinkt_go"

func main() {

	brightness := 15
	blinkt := NewBlinkt(brightness)

	checkPeriodSeconds := 60

	blinkt.SetClearOnExit(true)

	blinkt.Setup()

	Delay(100)

	for {
		num := getAstronautCount()

		r := 150
		g := 0
		b := 0
		blinkt.Clear()

		for pixel := 0; pixel < num; pixel++ {
			blinkt.SetPixel(pixel, r, g, b)
			blinkt.Show()
			Delay(100)
		}

		Delay(checkPeriodSeconds * 1000)
	}
	blinkt.Clear()
	blinkt.Show()
}
