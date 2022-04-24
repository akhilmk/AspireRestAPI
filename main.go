package main

import (
	"log"
	"net/http"
	"strings"

	api "github.com/akhilmk/AspireRestAPI/api"
	"github.com/gorilla/handlers"
)

func main() {
	log.Printf("Server Started !")

	router := api.NewRouter()
	headers := handlers.AllowedHeaders([]string{"Host", "Origin", "Connection", "Upgrade", "Sec-WebSocket-Key", "Sec-WebSocket-Version", "X-Requested-With", "Content-Type", "Authorization", "Accept", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "access-control-allow-origin", "access-control-allow-headers"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	credentials := handlers.AllowCredentials()

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins, credentials)(processRequestURL(router))))

}

// Middleware
func processRequestURL(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		// toke verification logic goes here.
		next.ServeHTTP(w, r)
	})
}
