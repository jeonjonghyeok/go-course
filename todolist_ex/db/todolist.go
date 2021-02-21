package db

import "github.com/learningspoons-go/todolist/todo"

func GetTodoLists() ([]todo.List, error) {
	rows, err := db.Query(`SELECT id, name FROM todo_list`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

}
