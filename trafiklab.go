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
	var f interface{}
    json.Unmarshal(body, &f)
	var r train
	v := m(m(m(m(f)["DPS"])["Trains"])["DpsTrain"].([]interface{})[0])
	r.Destination = v["Destination"].(string)
	r.LineNumber = int(v["LineNumber"].(float64))
	r.JourneyDirection = int(v["JourneyDirection"].(float64))
	r.TransportMode = v["TransportMode"].(string)
	return []train{r}
}

func m(j interface{}) map[string]interface{} {
	return j.(map[string]interface{})
}

func ParseString(body string) {
	fmt.Println(body)
}
