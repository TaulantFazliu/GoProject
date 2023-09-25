package routes

import (
	"log"
	"net/http"
)

func JSONLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("{\"method\": \"%s\", \"url\": \"%s\"}", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
