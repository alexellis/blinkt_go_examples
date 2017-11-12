package main

import (
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type cheerlight struct {
	Colour string `json:"field2"`
}

func getRGBFromColour(hexString string) (int, int, int) {
	//strip off the hash and decode the remaining hex value
	hexCol, _ := hex.DecodeString(strings.TrimPrefix(hexString, "#"))
	//return the ints representing rgb
	return int(hexCol[0]), int(hexCol[1]), int(hexCol[2])
}

func getCheerlightColours() (int, int, int) {
	var netClient = &http.Client{
		Timeout: time.Second * 3,
	}
	resp, getErr := netClient.Get("http://api.thingspeak.com/channels/1417/field/2/last.json")

	if getErr != nil {
		log.Panic(getErr)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		log.Panic(getErr)
	}

	result := cheerlight{}

	parseErr := json.Unmarshal(body, &result)
	if parseErr != nil {
		log.Panic("Can't parse response")
		return 0, 0, 0
	}

	return getRGBFromColour(result.Colour)

}
