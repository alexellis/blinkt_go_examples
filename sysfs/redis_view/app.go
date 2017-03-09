package main

import (
	"encoding/json"
	"fmt"
	"os"
)

import . "github.com/alexellis/blinkt_go/sysfs"

type LedColor struct {
	Red   int `json:"r"`
	Green int `json:"g"`
	Blue  int `json:"b"`
}

type LedMsg struct {
	Leds []LedColor `json:"leds"`
}

func newClient() *redis.Client {
	addr := os.Getenv("ADDR")
	client := redis.NewClient(&redis.Options{
		Addr:     addr + ":6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Output: PONG <nil>
	return client
}

func main() {
	brightness := 0.5
	blinkt := NewBlinkt(brightness)
	blinkt.SetClearOnExit(true)

	blinkt.Setup()

	client := newClient()
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	pubsub, err := client.Subscribe("lights")
	if err != nil {
		panic(err)
	}
	defer pubsub.Close()
	for {
		msg, err := pubsub.ReceiveMessage()
		if err != nil {
			panic(err)
		}
		if msg.Channel == "lights" {
			ledMsg := LedMsg{}
			jsonErr := json.Unmarshal([]byte(msg.Payload), &ledMsg)
			if jsonErr != nil {
				fmt.Println(jsonErr)
			}

			blinkt.Clear()

			var r, g, b int
			fmt.Println("Setting LEDs")

			for pixel := 0; pixel < 8; pixel++ {
				if pixel > len(ledMsg.Leds) {
					r = 0
					g = 0
					b = 0
				} else {
					r = ledMsg.Leds[pixel].Red
					g = ledMsg.Leds[pixel].Green
					b = ledMsg.Leds[pixel].Blue
				}
				blinkt.SetPixel(pixel, r, g, b)
			}
			blinkt.Show()
		}
	}
}
