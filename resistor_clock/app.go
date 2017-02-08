package main

import (
	"time"

	. "github.com/alexellis/blinkt_go"
)

func breakOut(colours []int) (int, int, int) {
	r := colours[0]
	g := colours[1]
	b := colours[2]
	return r, g, b
}

func main() {

	brightness := 25
	blinkt := NewBlinkt(brightness)

	blinkt.SetClearOnExit(true)

	colours := [10][3]int{
		{0, 0, 0},       //0 black
		{139, 69, 19},   //1 brown
		{255, 0, 0},     //2 red
		{255, 69, 0},    //3 orange
		{255, 255, 0},   //4 yellow
		{0, 255, 0},     //5 green
		{0, 0, 255},     //6 blue
		{128, 0, 128},   //7 violet
		{255, 255, 100}, //8 grey
		{255, 255, 255}, //9 white
	}

	blinkt.Setup()
	Delay(100)

	r := 0
	g := 0
	b := 0

	for {

		t := time.Now()
		hour := t.Hour()
		minute := t.Minute()

		hourten := hour / 10
		hourunit := hour % 10
		minuteten := minute / 10
		minuteunit := minute % 10

		r, g, b = breakOut(colours[hourten][:])
		blinkt.SetPixel(0, r, g, b)
		blinkt.SetPixel(1, r, g, b)

		r, g, b = breakOut(colours[hourunit][:])
		blinkt.SetPixel(2, r, g, b)
		blinkt.SetPixel(3, r, g, b)

		r, g, b = breakOut(colours[minuteten][:])
		blinkt.SetPixel(4, r, g, b)
		blinkt.SetPixel(5, r, g, b)

		r, g, b = breakOut(colours[minuteunit][:])
		blinkt.SetPixel(6, r, g, b)
		blinkt.SetPixel(7, r, g, b)

		blinkt.Show()
		Delay(500)

		blinkt.SetPixel(7, 0, 0, 0)

		blinkt.Show()
		Delay(500)

	}
}
