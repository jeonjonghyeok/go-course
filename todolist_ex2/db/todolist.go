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

func GetTodoList(list_id int) (todo.ListwithItem, error) {
	var list todo.ListwithItem
	rows, err := DB.Query(`SELECT l.ID, l.Name, i.ID, i.text, i.Done from todo_list l left join todo_item i on l.ID == i.todo_list_ID where l.ID == $1`, list_id)
	if err != nil {
		return list, err
	}
	defer rows.Close()
	list.Items = []todo.Item{}
	for rows.Next() {
		var (
			itemID   *int
			itemText *string
			itemDone *bool
		)
		if err := rows.Scan(&list.ID, &list.Name, &itemID, &itemText, &itemDone); err != nil {
			return list, err
		}
		if itemID != nil && itemText != nil && itemDone != nil {
			list.Items = append(list.Items, todo.Item{
				ID:   *itemID,
				Text: *itemText,
				Done: *itemDone,
			})
		}
	}
	return list, nil

}
