package log

import (
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newEntry.Debugf("[%s] - %s\n", r.Method, r.RequestURI)

		next.ServeHTTP(w, r)
	})
}
