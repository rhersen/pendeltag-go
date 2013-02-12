package main

import (
	"fmt"
	"strconv"
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

func p(p int) text {
	return text(strconv.Itoa(p) + " ")
}

func isPrime(n int) bool {
	if n < 4 {
		return true
	}

	for i := 2; i * i <= n; i++ {
		if n % i == 0 {
			return false
		}
	}

	return true
}

func primes(limit int) writer {
	children := []writer{}

	for i := 2; i < limit; i++ {
		if isPrime(i) {
			children = append(children, ec("span", p(i)))
		}
	}

	return element{"ol", children, a0()}
}

func index(w http.ResponseWriter, r *http.Request) {
	heading := text("This web app does not use any files.")
	fmt.Fprintf(w, "<!DOCTYPE html>")

	linkAttributes := map[string]string{
		"rel":"stylesheet",
		"href":"/css/"}

	ec("head", ec("title", heading), ea("link", linkAttributes)).html(w)
	ec("body", ec("h1", heading), primes(99999)).html(w)
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
