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

func CreateTodoList(name string) (list todo.List, err error) {
	list.Name = name
	err = DB.QueryRow(`INSERT INTO todo_list (name) VALUES ($1) RETURNING id`, name).Scan(&list.ID)
	return
}

func DeleteTodoList(id int) error {
	r, err := DB.Exec(`DELETE FROM todo_list where id=$1`, id)
	if err != nil {
		return err
	}
	if rn, err := r.RowsAffected(); rn == 0 || err != nil {
		return ErrorNotFound
	}
	return nil
}

func CreateTodoItem(listID int, text string, done bool) (item todo.Item, err error) {
	item.Text = text
	item.Done = done
	err = DB.QueryRow(`INSERT INTO todo_item (todo_list_id, text, done) VALUES ($1, $2, $3) 
		RETURNING id`, listID, text, done).Scan(&item.ID)
	return
}

func ModifyTodoList(listID int, name string) error {
	row, err := DB.Exec(`UPDATE todo_list SET name =$1 WHERE id = $2`, name, listID)
	if err != nil {
		return err
	}
	if r, err := row.RowsAffected(); r == 0 || err != nil {
		return ErrorNotFound
	}
	return nil
}

func GetTodoList(listID int) (todo.ListwithItem, error) {
	var list todo.ListwithItem
	rows, err := DB.Query(`SELECT l.id, l.name, i.id, i.text, i.done
		FROM todo_list l
		LEFT JOIN todo_item i
		ON l.id = i.todo_list_id
		WHERE l.id = $1`, listID)
	if err != nil {
		return list, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			id   int
			text string
			done bool
		)
		rows.Scan(&list.ID, &list.Name, &id, &text, &done)
		list.Items = append(list.Items, todo.Item{
			ID:   id,
			Text: text,
			Done: done,
		})
	}
	return list, nil
}
