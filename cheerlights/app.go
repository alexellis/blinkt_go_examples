package main

import (
	"os"
	"strconv"

	. "github.com/alexellis/blinkt_go/sysfs"
)

func getEnv(envVar string, assumed int) int {
	if value, exists := os.LookupEnv(envVar); exists {
		if period, err := strconv.Atoi(value); err == nil {
			return period
		}
	}
	return assumed
}

const defaultRefreshSeconds = 60
const envRefreshSeconds = "refresh_seconds"

func main() {

	brightness := 0.5
	blinkt := NewBlinkt(brightness)

	checkPeriodSeconds := getEnv(envRefreshSeconds, defaultRefreshSeconds)

	blinkt.SetClearOnExit(true)

	blinkt.Setup()

	Delay(100)

	for {

		r, g, b := getCheerlightColours()

		blinkt.Clear()
		blinkt.SetAll(r, g, b)
		blinkt.Show()
		Delay(checkPeriodSeconds * 1000)
	}
}
