package apis

import (
	"net/http"

	"github.com/Prateeknandle/url-shortener/handlers"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"POST",
		"/short", // will shorten the URL
		handlers.Urlshortner,
	},
	Route{
		"GET",
		"/{key}", // redirect to actual URL from shorten URL
		handlers.Redirecturl,
	},
}
