package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, `<!DOCTYPE html>` +
		`<html>` +
		`<head>` +
		`<title>The Stand Alone Web App</title>` +
		`<link rel="stylesheet" href="/css">` +
		`</head>` +
		`<body>` +
		`<h1>The Stand Alone Web App</h1>` +
		`</body>` +
		`</html>`)
}

func css(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	fmt.Fprintf(w, "body {background-color: #add8e6;}")
}

func main() {
    http.HandleFunc("/", index)
	http.HandleFunc("/css/", css)
    http.ListenAndServe(":8080", nil)
}
