package main

import (
	"net/http"
)

func main() {
    http.HandleFunc("/", index)
	http.HandleFunc("/css/", css)
	http.HandleFunc("/js/", js)
	http.HandleFunc("/station/", station)
	http.HandleFunc("/departures/", departures)
    http.ListenAndServe(":4000", nil)
}
