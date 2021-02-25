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
