package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type castWriter interface {
	Write(w http.ResponseWriter)
}

func handlePanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			log.Println("defer call")
			if r := recover(); r != nil {
				if e, ok := r.(castWriter); ok {
					e.Write(w)
				} else {
					internalError.Write(w)
				}
			}
		}()
		next.ServeHTTP(w, r)
	})
}

type simpleError struct {
	message string
	status  int
}

func (e simpleError) Write(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(e.message)
	log.Println(e.message)
}

var notFoundError = simpleError{
	message: "Not Found",
	status:  http.StatusNotFound,
}

var internalError = simpleError{
	message: "Internal Error",
	status:  http.StatusInternalServerError,
}

var malformedInputError = simpleError{
	message: "Malformed Input Error",
	status:  http.StatusBadRequest,
}
