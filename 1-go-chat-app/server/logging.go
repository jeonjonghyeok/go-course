package server

import (
	"log"
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (w responseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func loggingmiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writer := responseWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}
		log.Println(r.Method, r.RequestURI, writer.status)
		next.ServeHTTP(writer, r)
	})
}
