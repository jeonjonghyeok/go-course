package api

import (
	"net/http"

	"github.com/jeonjonghyeok/todolist_ex2/server"
)

var castWriter interface {
	Write(w http.ResponseWriter)
}

func handlePanic(w http.ResponseWriter, r *http.Request) {
	if r := recover(); r != nil {
		if e, ok := r.(castWriter); ok {
			e.Write(w)
			return
		}
	}
	defer r.Close()
	http.ServeHTTP(w, r)
}

type simpleError struct {
	message string
	status  int
}

func (e simpleError) Write(w http.ResponseWriter) {
	server.WriteHeader(e.status)
	http.NewEncoder(w).Encode(e.message)
}

var notFoundError = simpleError{
	message: "Not Found",
	status:  http.StatusBadRequest,
}

var internalError = simpleError{
	message: "Internal Error",
	status:  http.StatusInternalServerError,
}
