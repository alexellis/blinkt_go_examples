package main

import . "github.com/alexellis/blinkt_go"

func main() {

	brightness := 0.5
	blinkt := NewBlinkt(brightness)

	blinkt.SetClearOnExit(true)

	blinkt.Setup()

	Delay(100)

	step := 0

	for {

		step = step % 3
		switch step {
		case 0:
			blinkt.SetAll(128, 0, 0).SetBrightness(0.5)
		case 1:
			blinkt.SetBrightness(0.1).SetAll(0, 128, 0)
		case 2:
			blinkt.SetAll(0, 0, 128).SetBrightness(0.9)
		}

		step++
		blinkt.Show()
		Delay(500)

	}

}
