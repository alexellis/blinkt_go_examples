package main

import . "github.com/alexellis/blinkt_go"

func main() {

	brightness := 25
	blinkt := NewBlinkt(brightness)

	blinkt.SetClearOnExit(true)

	blinkt.Setup()

	Delay(100)

	step := 0

	for {

		step = step % 3
		switch step {
		case 0:
			blinkt.SetAll(128, 0, 0)
		case 1:
			blinkt.SetAll(0, 128, 0)
		case 2:
			blinkt.SetAll(0, 0, 128)
		}

		step++
		blinkt.Show()
		Delay(500)

	}

}
