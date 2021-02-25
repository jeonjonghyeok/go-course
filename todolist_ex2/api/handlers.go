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
