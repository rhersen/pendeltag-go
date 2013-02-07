package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	heading := text("The Stand Alone Web App")
	fmt.Fprintf(w, "<!DOCTYPE html>")
	element{"head", []writer{
			element{"title", []writer{heading}, map[string]string{}},
			element{"link", []writer{}, map[string]string{"rel":"stylesheet","href":"/css/"}}},
		map[string]string{}}.write(w)
	element{"body", []writer{element{"h1", []writer{heading}, map[string]string{}}}, map[string]string{}}.write(w)
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
