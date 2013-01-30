package main

import (
	"fmt"
	"encoding/json"
)

type train struct {
	Destination string
	LineNumber int
	JourneyDirection int
	TransportMode string
}

func Parse(body []byte) []train {
	r := []train{}
	for _, t := range getJsonTrains(body) {
		r = append(r, createTrain(t))
	}
	return r
}

func getJsonTrains(body []byte) []interface{} {
	var f interface{}
	json.Unmarshal(body, &f)
	return m(m(m(f)["DPS"])["Trains"])["DpsTrain"].([]interface{})
}

func createTrain(j interface{}) train {
	var r train
	v := m(j)
	r.Destination = v["Destination"].(string)
	r.LineNumber = int(v["LineNumber"].(float64))
	r.JourneyDirection = int(v["JourneyDirection"].(float64))
	r.TransportMode = v["TransportMode"].(string)
	return r
}

func m(j interface{}) map[string]interface{} {
	return j.(map[string]interface{})
}

func ParseString(body string) {
	fmt.Println(body)
}
