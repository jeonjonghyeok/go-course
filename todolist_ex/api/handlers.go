package api

import (
	"net/http"

	"github.com/learningspoons-go/todolist/db"
)

func getTodoLists(w http.ResponseWriter, r *http.Request) {
	lists, err := db.GetTodoList()
	must(err)
	writeJSON(w, lists)

}
