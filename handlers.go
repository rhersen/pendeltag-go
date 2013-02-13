package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, `<!DOCTYPE html>` +
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
	id := strings.Split(r.URL.Path, "/")[2]
	key := "get your own"

	resp, err := http.Get("https://api.trafiklab.se/sl/realtid/GetDpsDepartures.json?key=" + key + "&siteId=" + id + "&timeWindow=60")

	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("ERR")
		fmt.Println(err)
	} else {
		w.Write(ToJson(body))
	}

	resp.Body.Close()
}

func css(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	fmt.Fprintln(w, "body {margin: 5px; font-size: 22px;} section.table {width: 100%;} @media only screen and (orientation: landscape) {section.table {width: 50%;} body {font-size: 18px;}} section.table time {float: left; clear: left; width: 33.3%;} span.destination {float: left; width: 41.6%;} span.countdown {float: left; text-align: right; width: 25%;} .direction2 {background-color: #add8e6;} .direction1 {background-color: #e6adad;} @media only screen and (min-width: 800px) and (orientation: portrait) {body {font-size: 48px;}} @media only screen and (min-width: 800px) and (orientation: landscape) {body {font-size: 24px;}} body.pending {background: #d3d3d3;} div.departure {width: 100%;} a {text-decoration: none; font-size: 1.6em;} #expired {font-size: 0.8em;} #title, #predecessor, #successor {font-size: 1.4em;} #predecessor, #successor {color: #7F9F7F; font-style: italic;} #id {color: gray; font-style: italic; font-size: 0.5em;}")
}

func js(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	fmt.Fprintln(w, `function createStation(e){function t(){s.isPending()?$("body").addClass("pending"):$("body").removeClass("pending")}function n(n,r){function o(){s.setResponse(r),s.setUpdated(n.updated)}function u(){return _.first(_.first(n).Stops).SiteId-1}function a(){return _.first(_.first(n).Stops).SiteId+0}function f(){return _.first(_.first(n).Stops).SiteId+1}function l(){$("#title").html(names.abbreviate(_.first(_.first(n).Stops).StopAreaName)),$("#predecessor").html(u()),$("#successor").html(f()),$("#updated").html(n.updated)}function c(){$("section.table time").remove(),$("span.destination").remove(),$("span.countdown").remove(),_.each(n,h)}function h(e){var t="direction"+e.JourneyDirection,n=_.first(e.Stops).ExpectedDateTime,r=$(".table");$("<time></time>").appendTo(r).html(time.getTime(n)).addClass(t),$("<span></span>").appendTo(r).html(names.abbreviate(e.Destination)).addClass("destination").addClass(t),$("<span></span>").appendTo(r).addClass("countdown").addClass(t).data("time",_.first(e.Stops).ExpectedDateTime)}function p(){function t(e){return function(){i(e)}}var n=e?"touchend":"mouseup";$("#predecessor").bind(n,t(u())),$("#title").bind(n,t(a())),$("#successor").bind(n,t(f()))}o(),t(),l(),c(),p()}function r(t,n){function r(){function e(){var e=new Date;$("span.countdown").each(function(){var t=$(this).data("time");$(this).html(countdown.getCountdown(t,e))})}$("#expired").html(s.getDebugString()),e(),s.isExpired(new Date)&&i($("span#id").text())}$("span#id").text(t),e?$(".table").addClass("touch"):$(".table").addClass("mouse"),$("button.clear").click(function(){clearInterval(o)}),n&&(o=setInterval(r,n))}function i(e){s.setRequest((new Date).getTime()),t(),$("#title").unbind("mouseup touchend").html(e),$("#predecessor").unbind("mouseup touchend").html(" "),$("#successor").unbind("mouseup touchend").html(" "),$.ajax({url:"/departures/"+e,dataType:"json",cache:!1,success:function(e){n(e,(new Date).getTime())}}),$("span#id").text(e)}var s=expiry.create(),o;return{setResult:n,init:r}}function createCountdown(){function s(t){var n=t.getTimezoneOffset()*e;return(t.getTime()-n)%(i*r)}function o(n){var i=n.indexOf("T"),s=n.indexOf(":"),o=n.lastIndexOf(":");if(i<1||s<1||o<1)return undefined;var u=n.substring(i+1,s),a=n.substring(s+1,o),f=n.substring(o+1);return u*r+a*e+f*t}var e=6e4,t=1e3,n=60,r=n*e,i=24;return{getNow:s,millisSinceMidnight:o,getCountdown:function(r,i){function u(t){function r(e,t){var n=e%t;return(e-n)/t}function i(e){var t=e.toString();return t.length<2?"0"+t:t}var s=r(t,e)%n,o=r(t,1e3)%n,u=r(t,100)%10;return s+":"+i(o)+"."+u}var a=o(r)-s(i);return a<0?"-"+u(-a):u(a)}}}var expiry={};expiry.create=function(){function r(e){function n(e){return o(e)>30&&u(e.getTime())>20&&a(e.getTime())>10}return t===undefined||n(e)}function i(){return!e||e<t}function s(e){n=e}function o(e){return n?time.diff(countdown.getNow(e),countdown.millisSinceMidnight(n)):NaN}function u(e){return time.diff(e,t)}function a(t){return time.diff(t,e)}function f(e){t=e}function l(t){e=t}function c(){var e=new Date,t=o(e),n=u(e.getTime()),r=a(e.getTime());return t.toFixed(1)+"⊂"+n.toFixed(1)+"⊃"+r.toFixed(1)}var e,t,n;return{setUpdated:s,setRequest:f,setResponse:l,isExpired:r,isPending:i,getDebugString:c}};var names={abbreviate:function(e){function t(){return _.map([/^Upplands /,/^Stockholms /,/^T-/,/amn$/,/entrum$/],function(e){return{pattern:e,replacement:""}})}function n(e,t){return e.replace(t.pattern,t.replacement)}var r=[{pattern:/^Väster/,replacement:"V‧"},{pattern:/^Flemings/,replacement:"F‧"}];return _.reduce(r.concat(t()),n,e)}},countdown=createCountdown(),time={diff:function(t,n){return(t-n)/1e3},getTime:function(e){var t=/^.*T(.+):00$/.exec(e);if(t)return t[1];var n=/^.*T(.+)$/.exec(e);return n?n[1]:e}};`)
}

func station(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<!DOCTYPE html><html><head><meta content="true" name="HandheldFriendly"><meta content="width=device-width, height=device-height, user-scalable=no" name="viewport"><title>Station</title><link rel="stylesheet" href="/css"><script src="http://code.jquery.com/jquery-1.8.3.min.js"></script><script src="http://underscorejs.org/underscore-min.js"></script><script src="/js"></script></head><body><nav><span id="predecessor">9524</span><span id="title">Station</span><span id="successor">9526</span></nav><header><span id="id">9525</span><time id="updated">updated</time><button class="clear">⎚</button><span id="expired">?</span></header><section class="table"></section><script>createStation().init($('#id').text(), 256)</script></body></html>`)
}
