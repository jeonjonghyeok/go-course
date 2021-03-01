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

func createTodoList(w http.ResponseWriter, r *http.Request) {
	var req todo.List
	parseJSON(r.Body, &req)
	list, err := db.CreateTodoList(req.Name)
	must(err)
	writeJSON(w, list)
}

func deleteTodoList(w http.ResponseWriter, r *http.Request) {
	listID := parseIntParam(r, "list_id")
	must(db.DeleteTodoList(listID))
}

func createTodoItem(w http.ResponseWriter, r *http.Request) {
	var req todo.Item
	listID := parseIntParam(r, "list_id")
	parseJSON(r.Body, &req)
	item, err := db.CreateTodoItem(listID, req.Text, req.Done)
	must(err)
	writeJSON(w, item)

}

func modifyTodoList(w http.ResponseWriter, r *http.Request) {
	var req todo.List
	listID := parseIntParam(r, "list_id")
	parseJSON(r.Body, &req)
	must(db.ModifyTodoList(listID, req.Name))
}

func getTodoList(w http.ResponseWriter, r *http.Request) {
	listID := parseIntParam(r, "list_id")
	list, err := db.GetTodoList(listID)
	must(err)
	writeJSON(w, list)
}

func modifyTodoItem(w http.ResponseWriter, r *http.Request) {
	listID := parseIntParam(r, "list_id")
	itemID := parseIntParam(r, "item_id")
	var item todo.Item
	parseJSON(r.Body, &item)

	must(db.ModifyTodoItem(listID, itemID, item.Text, item.Done))

}

func deleteTodoItem(w http.ResponseWriter, r *http.Request) {
	listID := parseIntParam(r, "list_id")
	itemID := parseIntParam(r, "item_id")
	must(db.DeleteTodoItem(listID, itemID))
}
