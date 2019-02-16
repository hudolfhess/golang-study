package middlewares

import (
	"log"
	"net/http"
)

func ErrorMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("Internal server error: %s %s - Error: %s", r.Method, r.URL.Path, rec)
				http.Error(w, rec.(string), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	}
}
