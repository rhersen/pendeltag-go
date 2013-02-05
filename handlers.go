package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, `<!DOCTYPE html>` +
		`<html>` +
		`<head>` +
		`<meta content="true" name="HandheldFriendly">` +
		`<meta content="width=device-width, height=device-height, user-scalable=no" name="viewport">` +
		`<title>SL go</title>` +
		`<link rel="stylesheet" href="/css">`+
		//`<script src="http://code.jquery.com/jquery-1.8.3.min.js">`+
		//`</script>`+
		//`<script src="http://underscorejs.org/underscore-min.js">` +
		//`</script>`+
		//`<script src="/pendeltag-client.min.js">` +
		//`</script>` +
		`</head>` +
		`<body>` +
		`<h1>SL go</h1>` +
		`<ol>` +
		`<li>` +
		`<a href="station/9525">Tullinge</a>` +
		`</li>` +
		`<li>` +
		`<a href="station/9520">Södertälje</a>` +
		`</li>` +
		`<li>` +
		`<a href="station/9530">Södra</a>` +
		`</li>` +
		`<li>` +
		`<a href="station/9510">Karlberg</a>` +
		`</li>` +
		`<li>` +
		`<a href="station/9001">Centralen</a>` +
		`</li>` +
		`</ol>` +
		`</body>` +
		`</html>`)
}

func departures(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not yet implemented\n")
	fmt.Fprintf(w, "%s", r.URL.Path)
}

func css(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	fmt.Fprintf(w, "body {margin: 5px; font-size: 22px;} section.table {width: 100%;} @media only screen and (orientation: landscape) {section.table {width: 50%;} body {font-size: 18px;}} section.table time {float: left; clear: left; width: 33.3%;} span.destination {float: left; width: 41.6%;} span.countdown {float: left; text-align: right; width: 25%;} .direction2 {background-color: #add8e6;} .direction1 {background-color: #e6adad;} @media only screen and (min-width: 800px) and (orientation: portrait) {body {font-size: 48px;}} @media only screen and (min-width: 800px) and (orientation: landscape) {body {font-size: 24px;}} body.pending {background: #d3d3d3;} div.departure {width: 100%;} a {text-decoration: none; font-size: 1.6em;} #expired {font-size: 0.8em;} #title, #predecessor, #successor {font-size: 1.4em;} #predecessor, #successor {color: #7F9F7F; font-style: italic;} #id {color: gray; font-style: italic; font-size: 0.5em;}")
}

func station(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<!DOCTYPE html><html><head><meta content="true" name="HandheldFriendly"><meta content="width=device-width, height=device-height, user-scalable=no" name="viewport"><title>Station</title><link rel="stylesheet" href="/css"><script src="http://code.jquery.com/jquery-1.8.3.min.js"></script><script src="http://underscorejs.org/underscore-min.js"></script><script src="/pendeltag-client.min.js"></script></head><body><nav><span id="predecessor">9524</span><span id="title">Station</span><span id="successor">9526</span></nav><header><span id="id">9525</span><time id="updated">updated</time><button class="clear">⎚</button><span id="expired">?</span></header><section class="table"></section><script>createStation().init($('#id').text(), 256)</script></body></html>`)
	fmt.Fprintf(w, "%s", r.URL.Path)
}
