package main

import (
	"testing"
	"io/ioutil"
)

func TestShouldReturnTrains(t *testing.T) {
	ctx = t
	fixture, err := ioutil.ReadFile("huddinge.json")
	assert(err == nil)
	parsed := Parse(fixture)
	assertEqualsString("Södertälje centrum", parsed[0].Destination)
	assertEqualsString("Märsta", parsed[4].Destination)
}

func TestShouldReturnTrainSpecificFields(t *testing.T) {
	ctx = t
	fixture, err := ioutil.ReadFile("huddinge.json")
	assert(err == nil)
	train := Parse(fixture)[0]
	assertEqualsString("Södertälje centrum", train.Destination)
	assertEqualsInt(36, train.LineNumber)
	assertEqualsInt(1, train.JourneyDirection)
	assertEqualsString("TRAIN", train.TransportMode)
}

func TestShouldReturnStationSpecificFields(t *testing.T) {
	ctx = t
	fixture, err := ioutil.ReadFile("huddinge.json")
	assert(err == nil)
	stop := Parse(fixture)[0].Stops[0]
	assertEqualsString("0", stop.StopAreaNumber)
	assertEqualsString("Huddinge", stop.StopAreaName)
	assertEqualsString("2013-01-02T13:17:00", stop.TimeTabledDateTime)
	assertEqualsString("2013-01-02T13:17:19", stop.ExpectedDateTime)
	assertEqualsString("1 min", stop.DisplayTime)
}

func TestShouldHandleEmptyResponse(t *testing.T) {
	ctx = t
	fixture, err := ioutil.ReadFile("empty.json")
	assert(err == nil)
	assertEqualsInt(0, len(Parse(fixture)))
}
