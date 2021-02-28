package api

import (
	"net/http"

	"github.com/jeonjonghyeok/todolist_test/db"
)

func getTodoLists(w http.ResponseWriter, r *http.Request) {
	lists, err := db.GetTodoLists()
	must(err)
	writeJSON(w, lists)
}
