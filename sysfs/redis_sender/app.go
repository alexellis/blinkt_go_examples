package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/redis.v5"
	"os"
	"time"
)

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
	client := newClient()
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	i := 0
	dir := 0
	for {
		ledMsg := LedMsg{}
		ledMsg.Leds = make([]LedColor,8)
//		fmt.Printf("Dir=%d,i=%d\n",dir,i)
		ledMsg.Leds[i].Red = 255
		if(dir == 0) {
			i++
			if(i == 7) {
				dir = 1
			}
		} else if (dir == 1) {
			i--
			if i == 0 {
				dir = 0
			}
		}

		jsonValue, _:= json.Marshal(&ledMsg)

		pubErr := client.Publish("lights", string(jsonValue)).Err()
		if pubErr != nil {
			panic(pubErr)
		}
		time.Sleep(150 * time.Millisecond)
	}
}
