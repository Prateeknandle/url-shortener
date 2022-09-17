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
		"Post",
		"/short", // will shorten the url
		handlers.Urlshortner,
	},
}
