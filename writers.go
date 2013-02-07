package main

import (
	"fmt"
	"net/http"
)

type writer interface {
	write(http.ResponseWriter)
}

type text string

func (t text) write(w http.ResponseWriter) {
	fmt.Fprint(w, t)
}

type element struct {
	name text
	children []writer
	attributes map[string]string
}

func (e element) write(w http.ResponseWriter) {
	fmt.Fprintf(w, "<%s", e.name)

	for name, value := range e.attributes {
		fmt.Fprintf(w, " %s=%s", name, value)
	}

	fmt.Fprint(w, ">")

	for _, child := range e.children {
		child.write(w)
	}

	fmt.Fprintf(w, "</%s>", e.name)
}
