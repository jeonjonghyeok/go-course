package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func TodoListAPI() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/lists", getTodoLists).Methods(http.MethodGet)
	router.HandleFunc("/list", createTodoList).Methods(http.MethodPost)
	router.HandleFunc("/list/{list_id}", deleteTodoList).Methods(http.MethodDelete)
	router.HandleFunc("/list/{list_id}/item", createTodoItem).Methods(http.MethodPost)
	router.HandleFunc("/list/{list_id}", modifyTodoList).Methods(http.MethodPut)
	router.HandleFunc("/list/{list_id}", getTodoList).Methods(http.MethodGet)

	router.Use(handlePanic)

	return router
}
