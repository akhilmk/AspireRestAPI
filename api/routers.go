package api

import (
	"github.com/akhilmk/GoRESTAPI/constant"
	"github.com/akhilmk/GoRESTAPI/handlers"
	"github.com/akhilmk/GoRESTAPI/model"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	for _, route := range getRoutes() {
		handler := Logger(route.HandlerFunc, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

func getRoutes() []model.Route {
	return []model.Route{
		{
			Name:        "User",
			Method:      constant.GET,
			Pattern:     "/users",
			HandlerFunc: handlers.GetUsers,
		}, {
			Name:        "User",
			Method:      constant.GET,
			Pattern:     "/users/{id}",
			HandlerFunc: handlers.GetUser,
		}, {
			Name:        "User",
			Method:      constant.POST,
			Pattern:     "/users",
			HandlerFunc: handlers.AddUser,
		},
		{
			Name:        "User",
			Method:      constant.PUT,
			Pattern:     "/users/{id}",
			HandlerFunc: handlers.UpdateUser,
		},
		{
			Name:        "User",
			Method:      constant.DELETE,
			Pattern:     "/users/{id}",
			HandlerFunc: handlers.DeleteUser,
		},
	}
}
