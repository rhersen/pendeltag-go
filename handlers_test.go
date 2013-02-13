package main

import (
	"strings"
	"testing"
	"net/url"
	"net/http"
)

type ResponseWriterMock struct {
	written string
	header http.Header
}

func (mock *ResponseWriterMock) getWritten() string {
	return string(mock.written)
}

func (mock *ResponseWriterMock) Header() http.Header {
	if len(mock.header) == 0 {
		mock.header = make(http.Header)
	}

	return mock.header
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

func TestCss(t *testing.T) {
	ctx = t
	result := new(ResponseWriterMock)
	request := createRequest();
	assertEqualsString("", result.Header().Get("Content-Type"))
	css(result, request)
	assert(strings.Contains(result.getWritten(), "section.table"))
	assertEqualsString("text/css", result.Header().Get("Content-Type"))
}

func TestStationShouldContainTable(t *testing.T) {
	ctx = t
	result := new(ResponseWriterMock)
	request := createRequest();
	station(result, request)
	assert(strings.Contains(result.getWritten(), `<section class="table">`))
}

func TestJsShouldNotContainTable(t *testing.T) {
	ctx = t
	result := new(ResponseWriterMock)
	request := createRequest();
	js(result, request)
	assert(!strings.Contains(result.getWritten(), `<section class="table">`))
	assertEqualsString("application/javascript", result.Header().Get("Content-Type"))
}

func createRequest() *http.Request {
	r := new(http.Request)
	url := new(url.URL)
	r.URL = url
	return r
}