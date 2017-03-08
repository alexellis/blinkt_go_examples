package main

import . "github.com/alexellis/blinkt_go/sysfs"

func main() {

	brightness := 0.5
	blinkt := NewBlinkt(brightness)

	checkPeriodSeconds := 60

	blinkt.SetClearOnExit(true)

	blinkt.Setup()

	Delay(100)

	for {

		r, g, b := getCheerlightColours()

		blinkt.Clear()
		blinkt.SetAll(r, g, b)
		blinkt.Show()
		Delay(checkPeriodSeconds * 1000)
	}
}
