package server

import (
	"log"
	"net/http"
)

type statusCaptureResponseWriter struct {
	http.ResponseWriter
	status int
}

func (s *statusCaptureResponseWriter) WriteHeader(statusCode int) {
	s.status = statusCode
	s.ResponseWriter.WriteHeader(statusCode)
}

func LoggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusCapturingWriter := &statusCaptureResponseWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}
		next.ServeHTTP(statusCapturingWriter, r)

		log.Printf("[%s] %s - %d\n", r.Method, r.RequestURI, statusCapturingWriter.status)
	})
}
