package server

import (
	"log"
	"net/http"
)

type responseWriterCapture struct {
	http.ResponseWriter
	status int
}

func (r *responseWriterCapture) WriteHeader(status int) {
	r.status = status
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resWriter := &responseWriterCapture{
			ResponseWriter: w,
			status:         http.StatusOK,
		}
		next.ServeHTTP(resWriter, r)
		log.Println(r.Method, r.RequestURI, resWriter.status)
	})
}
