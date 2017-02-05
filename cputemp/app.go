package main

import (
	. "github.com/alexellis/blinkt_go"
	"os/exec"
	"strings"
	"bytes"
	"fmt"
	"strconv"
)

func getTemperature() float64 {
	targetCmd := exec.Command("vcgencmd", "measure_temp")
	var out bytes.Buffer
	targetCmd.Stdout = &out
	err := targetCmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	temp := out.String()

	//temp=35.8'C
	tempVal := temp[strings.Index(temp, "=") + 1 : len(temp) - 3 ]
	celcius, _ := strconv.ParseFloat(tempVal, 64)
	return celcius
}

func min(x float64, y float64) float64 {
    if x < y {
        return x
    }
    return y
}

func showGraph(bl *Blinkt, v float64, r int, g int, b int) {
	v = v * 8
	one:=float64(1.0)

	for i:=0; i< 8; i++ {
		if v < 0 {
			r, g, b = 0, 0, 0
		} else {
    			r = int(float64(r) * min(v, one))
    			g = int(float64(g) * min(v, one))
    			b = int(float64(b) * min(v, one))
		}
		bl.SetPixel(i, r, g, b)
		v = v - 1
	}
	bl.Show()
}

func main() {
	blinkt := NewBlinkt(20)
	blinkt.Setup()

	celcius := getTemperature()
	fmt.Println(celcius)

	v := celcius / 100
	showGraph(&blinkt, v, 255, 255, 255)
}
