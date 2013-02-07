package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func a0() map[string]string {
	return map[string]string{}
}

func ec(name text, child ...writer) writer {
	return element{name, child, a0()}
}

func ea(attributes map[string]string) writer {
	return element{"link", []writer{}, attributes}
}

func index(w http.ResponseWriter, r *http.Request) {
	heading := text("The Stand Alone Web App")
	fmt.Fprintf(w, "<!DOCTYPE html>")

	linkAttributes := map[string]string{
		"rel":"stylesheet",
		"href":"/css/"}

	ec("head", ec("title", heading), ea(linkAttributes)).write(w)
	ec("body", ec("h1", heading)).write(w)
}

func css(w http.ResponseWriter, req *http.Request) {
	r := rand.Int() & 255
	g := rand.Int() & 255
	b := rand.Int() & 255
	w.Header().Set("Content-Type", "text/css")
	fmt.Fprintf(w,
		"body {background-color:rgb(%d,%d,%d);color:rgb(%d,%d,%d);}",
		r, g, b,
		g, b, r)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/css/", css)
	http.ListenAndServe(":8080", nil)
}
