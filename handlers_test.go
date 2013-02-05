package main

import (
	"strings"
	"testing"
	"net/url"
	"net/http"
)

type ResponseWriterMock struct {
	written string
}

func (mock *ResponseWriterMock) getWritten() string {
	return string(mock.written)
}

func (mock *ResponseWriterMock) Header() http.Header {
	var r http.Header
	return r
}

func (mock *ResponseWriterMock) Write(response []byte) (int, error) {
	mock.written += string(response)
	return 0, nil
}

func (mock *ResponseWriterMock) WriteHeader(int) {}

func TestIndexShouldContainLinkToStation(t *testing.T) {
	ctx = t
	result := new(ResponseWriterMock)
	request := createRequest();
	index(result, request)
	assert(strings.Contains(result.getWritten(), `<a href="station/`))
}

func TestDeparturesIsNotImplemented(t *testing.T) {
	ctx = t
	result := new(ResponseWriterMock)
	request := createRequest();
	departures(result, request)
	assert(strings.Contains(result.getWritten(), "not yet"))
}

func TestStationShouldContainTable(t *testing.T) {
	ctx = t
	result := new(ResponseWriterMock)
	request := createRequest();
	station(result, request)
	assert(strings.Contains(result.getWritten(), `<section class="table">`))
}

func createRequest() *http.Request {
	r := new(http.Request)
	url := new(url.URL)
	r.URL = url
	return r
}