package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
	resp, err := http.Get("https://api.trafiklab.se/sl/realtid/GetDpsDepartures.json?key=8589732b19f6c9a78b004aff74f28d98&siteId=9525&timeWindow=60")

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
