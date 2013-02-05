package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
    http.HandleFunc("/", index)
	http.HandleFunc("/css/", css)
	http.HandleFunc("/station/", station)
	http.HandleFunc("/departures/", departures)
    http.ListenAndServe(":8080", nil)
}

func main0() {
	key := "8589732b19f6c9a78b004aff74f28d98"
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
		parsed := Parse(body)
		fmt.Println(parsed)
	}

	resp.Body.Close()
}
