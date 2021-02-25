package api

import (
	"net/http"

	"github.com/jeonjonghyeok/todolist_ex2/db"
)

func getTodoLists(w http.ResponseWriter, r *http.Request) {
	lists, err := db.GetTodoLists()
	must(err)
	writeJSON(w, lists)
}

func getTodoList(w http.ResponseWriter, r *http.Request) {
	list_id := parseIntParam(r, "list_id")
	list, err := db.GetTodoList(list_id)
	must(err)
	writeJSON(w, list)
}
