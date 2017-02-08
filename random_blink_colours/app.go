package main

import (
	"math/rand"

	. "github.com/alexellis/blinkt_go"
)

func main() {

	brightness := 25
	blinkt := NewBlinkt(brightness)

	blinkt.SetClearOnExit(true)

	blinkt.Setup()

	Delay(100)

	for {

		for pixel := 0; pixel < 8; pixel++ {
			blinkt.SetPixel(pixel, rand.Intn(255), rand.Intn(255), rand.Intn(255))
		}

		blinkt.Show()
		Delay(50)
	}

}
