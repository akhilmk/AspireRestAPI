package api

import (
	"net/http"

	handlers "github.com/akhilmk/AspireRestAPI/handlers"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	for _, route := range routes {
		handler := Logger(route.HandlerFunc, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

const (
	POST = "POST"
	GET  = "GET"
)

type Routes []Route

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routes = Routes{
	Route{
		"Loan",
		GET,
		"/loan",
		handlers.GetLoan,
	},
	Route{
		"Loan",
		POST,
		"/loan",
		handlers.AddLoan,
	},
}
