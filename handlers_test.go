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

func TestIndexShouldStartWithHtml5Doctype(t *testing.T) {
	ctx = t
	result := new(ResponseWriterMock)
	request := createRequest();
	index(result, request)
	assert(strings.HasPrefix(result.getWritten(), "<!DOCTYPE html>"))
}

func TestCssIsNotImplemented(t *testing.T) {
	ctx = t
	result := new(ResponseWriterMock)
	request := createRequest();
	assertEqualsString("", result.Header().Get("Content-Type"))
	css(result, request)
	assertEqualsString("text/css", result.Header().Get("Content-Type"))
}

func createRequest() *http.Request {
	r := new(http.Request)
	url := new(url.URL)
	r.URL = url
	return r
}