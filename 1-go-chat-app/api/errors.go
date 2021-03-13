package api

import (
	"encoding/json"
	"net/http"
)

func must(e error) {
	panic(internalError)
}

type apiErrorWriter interface {
	Write(w http.ResponseWriter)
}

func handlePanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				if c, ok := r.(apiErrorWriter); ok {
					c.Write(w)
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
	w.WriteHeader(e.status)
	json.NewEncoder(w).Encode(e.message)
}

var notFoundError = simpleError{
	message: "Not found",
	status:  http.StatusNotFound,
}

var internalError = simpleError{
	message: "Internal error",
	status:  http.StatusInternalServerError,
}

var malformedInputError = simpleError{
	message: "Malformed input",
	status:  http.StatusBadRequest,
}
