package db

import (
	"github.com/jeonjonghyeok/todolist_ex2/todo"
)

func GetTodoLists() ([]todo.List, error) {
	lists := []todo.List{}
	rows, err := DB.Query(`SELECT ID, Name FROM todo_list`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var list todo.List
		err := rows.Scan(&list.ID, &list.Name)
		if err != nil {
			return lists, err
		}
		lists = append(lists, list)
	}
	return lists, nil
}

func GetTodoList(todoListID int) (todo.ListwithItem, error) {
	var list todo.ListwithItem
	rows, err := DB.Query(`SELECT l.id, l.name, i.id, i.text, i.done 
		FROM todo_list l
		LEFT JOIN todo_item i ON l.id = i.todo_list_id 
		WHERE l.id = $1`, todoListID)
	if err != nil {
		return list, err
	}
	defer rows.Close()
	list.Items = []todo.Item{}
	var gotTodoList bool
	for rows.Next() {
		var (
			itemID   *int
			itemText *string
			itemDone *bool
		)
		if err := rows.Scan(&list.ID, &list.Name, &itemID, &itemText, &itemDone); err != nil {
			return list, err
		}
		gotTodoList = true
		if itemID != nil && itemText != nil && itemDone != nil {
			list.Items = append(list.Items, todo.Item{
				ID:   *itemID,
				Text: *itemText,
				Done: *itemDone,
			})
		}
		if !gotTodoList {
			return list, ErrorNotFound
		}
	}
	return list, nil

}
