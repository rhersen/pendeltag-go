package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	key := "get your own at trafiklab.se"
	resp, err := http.Get("https://api.trafiklab.se/sl/realtid/GetDpsDepartures.json?key=" + key + "&siteId=9525&timeWindow=60")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	} else {
		Parse(body)
	}

	resp.Body.Close()
}
