package server

import "net/http"

type responseWriterCapture struct {
	http.ResponseWriter
	status int
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resWriter := &responseWriterCapture{
			ResponseWriter: w,
			status:         http.StatusOK,
		}
		next.ServeHTTP(resWriter, r)
	})
}
