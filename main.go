package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<!DOCTYPE html>` +
		`<html>` +
		`<head>` +
		`<title>The Stand Alone Web App</title>` +
		`<link rel="stylesheet" href="/css/">` +
		`</head>` +
		`</html>`)
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
