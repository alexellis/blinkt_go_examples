package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	. "github.com/alexellis/blinkt_go"
)

func Delay(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func main() {

	brightness := 25
	blinkt := NewBlinkt(brightness)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	go func() {
		for range signalChan {
			fmt.Println("Control + C")
			blinkt.Clear()
			blinkt.Show()
			os.Exit(1)
		}
	}()

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
