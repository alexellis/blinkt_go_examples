package main

import "encoding/json"
import "net/http"
import "io/ioutil"
import "log"
import "time"

type People struct {
  Number	 int `json:"number"`

}

func main() {
	var netClient = &http.Client{
		Timeout: time.Second * 3,
	}
	resp, getErr := netClient.Get("http://api.open-notify.org/astros.json")
	if getErr != nil {
		log.Panic(getErr)
	}
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Panic(getErr)
	}

	var people People
	parseErr := json.Unmarshal(body, &people)
	if parseErr !=nil {
		log.Panic("Can't parse response")
		return
	}
	log.Println(string(body))
	log.Printf("People: %d", people.Number)
}
