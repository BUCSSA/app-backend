package app_backend

import (
	"github.com/gorilla/mux"
)

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		IndexHandler,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
