package main

import (
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	handler := http.FileServer(http.Dir("C:/Users/ybjqx/Music/"))
	http.Handle("/", http.StripPrefix("/", handler))
	http.ListenAndServe(":8081", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(handler))
}
