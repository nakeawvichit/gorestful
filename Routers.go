package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"index",
		"GET",
		"/",
		index,
	},
	Route{
		"say",
		"GET",
		"/say/{name}",
		say,
	}, Route{
		"api",
		"GET",
		"/api/",
		calljson,
	}, Route{
		"save",
		"POST",
		"/api/",
		save,
	},
}
