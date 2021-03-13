package db

import "github.com/jeonjonghyeok/go-run/1-go-chat-app/chat"

func GetRooms() ([]chat.Room, error) {
	rooms := []chat.Room{}
	rows, err := db.Query(`SELECT id, name from chatrooms`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var room chat.Room
		if err := rows.Scan(&room.ID, &room.Name); err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}
	return rooms, nil
}

func CreateRoom(name string) (id int, err error) {
	if err = db.QueryRow(`INSERT INTO chatrooms (name) VALUES ($1) RETURNING id`, name).Scan(&id); err != nil {
		return 0, err
	}
	return
}
