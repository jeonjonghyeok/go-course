package api

import (
	"encoding/json"
	"net/http"
)

type castWriter interface {
	Write(w http.ResponseWriter)
}

func handlePanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				if e, ok := r.(castWriter); ok {
					e.Write(w)
				}
			}
		}()
	})
}

type simpleError struct {
	message string
	status  int
}

func (e simpleError) Write(w http.ResponseWriter) {
	w.WriteHeader(e.status)
	json.NewEncoder(w).Encode(e.message)
}

var notFoundError = simpleError{
	message: "Not Found",
	status:  http.StatusBadRequest,
}

var internalError = simpleError{
	message: "Internal Error",
	status:  http.StatusInternalServerError,
}
