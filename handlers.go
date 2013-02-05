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
	fmt.Fprintf(w, "not yet implemented\n")
	fmt.Fprintf(w, "%s", r.URL.Path)
}

func station(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<!DOCTYPE html><html><head><meta content="true" name="HandheldFriendly"><meta content="width=device-width, height=device-height, user-scalable=no" name="viewport"><title>Station</title><link rel="stylesheet" href="/css"><script src="http://code.jquery.com/jquery-1.8.3.min.js"></script><script src="http://underscorejs.org/underscore-min.js"></script><script src="/pendeltag-client.min.js"></script></head><body><nav><span id="predecessor">9524</span><span id="title">Station</span><span id="successor">9526</span></nav><header><span id="id">9525</span><time id="updated">updated</time><button class="clear">⎚</button><span id="expired">?</span></header><section class="table"></section><script>createStation().init($('#id').text(), 256)</script></body></html>`)
	fmt.Fprintf(w, "%s", r.URL.Path)
}
