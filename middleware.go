package contactqr

import (
	"log"
	"net/http"
)

// LoggingMiddleware logs all requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != `/favicon.ico` {
			log.Println(r.Method, r.RequestURI)
		}
		next.ServeHTTP(w, r)
	})
}
