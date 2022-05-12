package api

import (
	"log"
	"net/http"
	"strings"

	"github.com/akhilmk/GoRESTAPI/constant"
	"github.com/akhilmk/GoRESTAPI/handlers"
	"github.com/akhilmk/GoRESTAPI/model"
	goHandler "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func StartService() {
	router := newRouter()
	headers := goHandler.AllowedHeaders([]string{"Host", "Origin", "Connection", "Upgrade", "Sec-WebSocket-Key", "Sec-WebSocket-Version", "X-Requested-With", "Content-Type", "Authorization", "Accept", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "access-control-allow-origin", "access-control-allow-headers"})
	methods := goHandler.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"})
	origins := goHandler.AllowedOrigins([]string{"*"})
	credentials := goHandler.AllowCredentials()

	log.Fatal(http.ListenAndServe(":8080", goHandler.CORS(headers, methods, origins, credentials)(processRequestURL(router))))
}

func newRouter() *mux.Router {
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

// Middleware
func processRequestURL(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		// token verification logic goes here.
		next.ServeHTTP(w, r)
	})
}
