package main

import (
	"fmt"
	"math/rand"
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
	child writer
}

func (e element) write(w http.ResponseWriter) {
	fmt.Fprintf(w, "<%s>", e.name)
	e.child.write(w)
	fmt.Fprintf(w, "</%s>", e.name)
}

func index(w http.ResponseWriter, r *http.Request) {
	heading := text("The Stand Alone Web App")
	fmt.Fprintf(w, "<!DOCTYPE html>")
	element{"head", element{"title", heading}}.write(w)
	element{"body", element{"h1", heading}}.write(w)
}

func css(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	fmt.Fprintf(w, "body {background-color:rgb(%d,%d,%d);}",
		rand.Int() & 255, rand.Int() & 255, rand.Int() & 255)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/css/", css)
	http.ListenAndServe(":8080", nil)
}
