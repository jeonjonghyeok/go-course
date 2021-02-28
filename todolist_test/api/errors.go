package api

import (
	"net/http"
)

type responseWriter struct {
	Write (http.ResponseWriter)
}

func errorPanic(err error) {
	if err != nil {
		if r := recover(); r != nil {
			r.(responseWriter)
		}
	}

}

type simpleError struct {
	message string
	status  int
}
type errorResponse struct {
	Error string `json:"error"`
}

/*
func (e simpleError) Write(w http.ResponseWriter) {
	w.WriteHeader(e.status)
	json.NewEncoder(w).Encode(errorResponse{e.message})
}
*/

var notFoundError = simpleError{
	message: "Not Found",
	status:  http.StatusNotFound,
}

var internalError = simpleError{
	message: "Internal Server Error",
	status:  http.StatusInternalServerError,
}
