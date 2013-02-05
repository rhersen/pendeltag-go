package main

import (
	"fmt"
	"runtime"
	"testing"
)

var ctx *testing.T

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
