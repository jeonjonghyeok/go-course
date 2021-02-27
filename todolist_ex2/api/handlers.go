package api

import (
	"net/http"

	"github.com/jeonjonghyeok/todolist_ex2/db"
	"github.com/jeonjonghyeok/todolist_ex2/todo"
)

func getTodoLists(w http.ResponseWriter, r *http.Request) {
	lists, err := db.GetTodoLists()
	must(err)
	writeJSON(w, lists)
}

func getTodoList(w http.ResponseWriter, r *http.Request) {
	listID := parseIntParam(r, "list_id")
	list, err := db.GetTodoList(listID)
	must(err)
	writeJSON(w, list)
}

func createTodoList(w http.ResponseWriter, r *http.Request) {
	var req todo.List
	parseJSON(r.Body, &req)
	list, err := db.CreateTodoList(req.Name)
	must(err)
	writeJSON(w, list)
}
