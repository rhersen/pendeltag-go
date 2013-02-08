package main

import (
	"fmt"
	"net/http"
)

type writer interface {
	html(http.ResponseWriter)
	css(http.ResponseWriter)
}

type text string

func (t text) css(w http.ResponseWriter) {
	fmt.Fprint(w, t)
}

func (t text) html(w http.ResponseWriter) {
	fmt.Fprint(w, t)
}

type element struct {
	name string
	children []writer
	attributes map[string]string
}

func (e element) css(w http.ResponseWriter) {
	fmt.Fprintf(w, "%s {", e.name)

	for name, value := range e.attributes {
		fmt.Fprintf(w, "%s:%s;", name, value)
	}

	fmt.Fprintf(w, "}")
}

func (e element) html(w http.ResponseWriter) {
	fmt.Fprintf(w, "<%s", e.name)

	for name, value := range e.attributes {
		fmt.Fprintf(w, " %s=%s", name, value)
	}

	fmt.Fprint(w, ">")

	for _, child := range e.children {
		child.html(w)
	}

	fmt.Fprintf(w, "</%s>", e.name)
}
