package routes

import (
	"net/http"

	"github.com/yinichai/coconut/handlers"
)

const (
	version string = "1.0.0"
)

// Route endpoint model
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes array of Route
type Routes []Route

var routes = Routes{
	Route{
		"ping",
		"GET",
		"/_ping",
		handlers.Ping(),
	},

	Route{
		"search item",
		"GET",
		"/search",
		handlers.Search,
	},
}
