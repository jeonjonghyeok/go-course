package db

import (
	"github.com/jeonjonghyeok/todolist_test/todo"
)

func GetTodoLists() ([]todo.List, error) {
	rows, err := db.Query(`select id, name from todo_list`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lists := []todo.List{}
	for rows.Next() {
		var list todo.List
		if err := rows.Scan(&list.ID, &list.Name); err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}
	return lists, nil
}
