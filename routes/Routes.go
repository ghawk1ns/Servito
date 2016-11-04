package routes

import (
	"net/http"
	"github.com/ghawk1ns/servito/handlers"
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
		handlers.Index,
	},
	Route{
		"Health",
		"GET",
		"/health",
		handlers.Health,
	},
}