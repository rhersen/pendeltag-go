package main

import (
	"encoding/json"
)

type train struct {
	Destination string
	LineNumber int
	JourneyDirection int
	TransportMode string
	Stops []stop
}

type stop struct {
	StopAreaNumber string
	StopAreaName string
	TimeTabledDateTime string
	ExpectedDateTime string
	DisplayTime string
}

func ToJson(body []byte) []byte {
	r, _ := json.Marshal(Parse(body))
	return r
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
	trains := m(m(f)["DPS"])["Trains"]

	if trains == nil {
		return nil
	}

	return m(trains)["DpsTrain"].([]interface{})
}

func createTrain(j interface{}) train {
	var r train
	v := m(j)
	r.Destination = v["Destination"].(string)
	r.LineNumber = int(v["LineNumber"].(float64))
	r.JourneyDirection = int(v["JourneyDirection"].(float64))
	r.TransportMode = v["TransportMode"].(string)
	r.Stops = []stop{createStop(v)}
	return r
}

func createStop(v map[string]interface{}) stop {
	var s stop
	s.StopAreaNumber = v["StopAreaNumber"].(string)
	s.StopAreaName = v["StopAreaName"].(string)
	s.TimeTabledDateTime = v["TimeTabledDateTime"].(string)
	s.ExpectedDateTime = v["ExpectedDateTime"].(string)
	s.DisplayTime = v["DisplayTime"].(string)
	return s
}

func m(j interface{}) map[string]interface{} {
	return j.(map[string]interface{})
}
