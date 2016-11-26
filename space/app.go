package main

import . "github.com/alexellis/blinkt_go"
import "time"
import "os/signal"
import "os"
import "fmt"

func Delay(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

// Fill in the gaps here, see the jsonclient example
func getSpace() int {
	return 6
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

		num := getSpace()

		r := 150
		g := 0
		b := 0
		blinkt.Clear()

		for pixel := 0; pixel < num; pixel++ {
			blinkt.SetPixel(pixel, r, g, b)
			blinkt.Show()
			Delay(100)
		}

		Delay(5000)
	}
	blinkt.Clear()
	blinkt.Show()
}
