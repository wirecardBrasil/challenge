package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Client",
		"GET",
		"/client",
		ConsultAllClients,
	},
	Route{
		"Client",
		"POST",
		"/client",
		InsertClient,
	},
}
