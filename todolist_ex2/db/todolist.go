package db

func GetTodoLists() ([]todo.List, error) {
	rows, err := db.Query(`SELECT l.id, l.name, i.id, i.text FROM todo_list l left join item_list i on i.ID = l.ID`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {

	}

}
