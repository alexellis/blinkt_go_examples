package main

import . "github.com/alexellis/blinkt_go"

func main() {
	brightness := 0.5
	blinkt := NewBlinkt(brightness)

	blinkt.SetClearOnExit(true)

	blinkt.Setup()

	Delay(100)

	r := 150
	g := 0
	b := 0
	for {
		for pixel := 0; pixel < 8; pixel++ {
			blinkt.Clear()
			blinkt.SetPixel(pixel, r, g, b)
			blinkt.Show()
			Delay(100)
		}
		for pixel := 7; pixel > 0; pixel-- {
			blinkt.Clear()
			blinkt.SetPixel(pixel, r, g, b)
			blinkt.Show()
			Delay(100)
		}
	}
	blinkt.Clear()
	blinkt.Show()
}
