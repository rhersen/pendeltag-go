package main

import (
	"fmt"
	"runtime"
	"testing"
	"io/ioutil"
)

var ctx *testing.T

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

func assert(b bool) {
	if !b {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d:\n", file, line)
		ctx.Errorf("%s:%d\n", file, line)
	}
}

func assertArrayEquals(expected []int64, actual []int64) {
	if len(expected) != len(actual) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: expected %d but got %d\n", file, line, expected, actual)
		ctx.Errorf("%s:%d expected %d but got %d\n", file, line, expected, actual)
		return
	}

	if len(expected) != len(actual) {
		_, file, line, _ := runtime.Caller(1)
		ctx.Errorf("%s:%d expected %d but got %d\n", file, line, expected, actual)
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			_, file, line, _ := runtime.Caller(1)
			fmt.Printf("%s:%d: expected %d but got %d\n", file, line, expected, actual)
			ctx.Errorf("%s:%d expected %d but got %d\n", file, line, expected, actual)
		}
	}
}

func assertEqualsInt(expected int, actual int) {
	if expected != actual {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: expected %d but got %d\n", file, line, expected, actual)
		ctx.Errorf("expected %d but got %d\n", expected, actual)
	}
}

func assertEqualsString(expected string, actual string) {
	if expected != actual {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: expected '%s' but got '%s'\n", file, line, expected, actual)
		ctx.Errorf("expected %s but got %s\n", expected, actual)
	}
}

func assertEqualsFloat(expected float64, actual float64) {
	if expected != actual {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: expected %f but got %f\n", file, line, expected, actual)
		ctx.Errorf("expected %f but got %f\n", expected, actual)
	}
}
