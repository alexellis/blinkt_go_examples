package main

import (
	"math/rand"
	"time"

	. "github.com/alexellis/blinkt_go"
)

func shuffleAndSlice(arr []int) []int {

	t := time.Now()
	rand.Seed(int64(t.Nanosecond()))

	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i)
		arr[i], arr[j] = arr[j], arr[i]
	}

	subsetSize := rand.Intn(5) + 1 // +1 as zero based
	return arr[:subsetSize]
}

func isIn(s *[]int, e *int) bool {
	for _, a := range *s {
		if a == *e {
			return true
		}
	}
	return false
}

func main() {

	brightness := 0.5
	blinkt := NewBlinkt(brightness)

	blinkt.SetClearOnExit(true)

	blinkt.Setup()

	Delay(100)

	nums := []int{0, 1, 2, 3, 4, 5, 6, 7}

	for {

		//There must be a more elegant way of doing this
		pixels := shuffleAndSlice(nums)
		for _, i := range nums {
			if isIn(&pixels, &i) {
				blinkt.SetPixel(i, 255, 150, 0)
			} else {
				blinkt.SetPixel(i, 0, 0, 0)
			}
		}

		blinkt.Show()
		Delay(50)
	}

}
