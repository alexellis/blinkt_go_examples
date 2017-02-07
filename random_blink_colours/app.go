package main

import (
	"fmt"
	"math/rand"
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

	for {

		for pixel := 0; pixel < 8; pixel++ {
			blinkt.SetPixel(pixel, rand.Intn(255), rand.Intn(255), rand.Intn(255))
		}

		blinkt.Show()
		Delay(50)
	}

}
