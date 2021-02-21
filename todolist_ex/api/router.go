package api

import (
	"net/http"

	"github.com/golila/mux"
)

func TodoListAPI() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/todos", getTodoList).Methods(http.MethodGet)
	router.HandleFunc("/todos", createTodoList)Methods(http.MethodPost)
	router.use(handlePanic)

	return router

}
