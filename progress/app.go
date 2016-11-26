package main

import . "github.com/alexellis/blinkt_go"
import "time"
import "os/signal"
import "os"
import "fmt"

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
