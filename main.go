package main

import (
	"fmt"
	"net/http"
)

func a0() map[string]string {
	return map[string]string{}
}

func ec(name string, child ...writer) writer {
	return element{name, child, a0()}
}

func ea(name string, attributes map[string]string) writer {
	return element{name, []writer{}, attributes}
}

func index(w http.ResponseWriter, r *http.Request) {
	heading := text("This web app does not use any files.")
	fmt.Fprintf(w, "<!DOCTYPE html>")

	linkAttributes := map[string]string{
		"rel":"stylesheet",
		"href":"/css/"}

	ec("head", ec("title", heading), ea("link", linkAttributes)).html(w)
	ec("body", ec("h1", heading)).html(w)
}

func css(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/css")

	properties := map[string]string{
		"background-color":"darkslategray",
		"color":"ivory"}

	ea("body", properties).css(w)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/css/", css)
	http.ListenAndServe(":8080", nil)
}
