package db

import "github.com/jeonjonghyeok/go-run/1-go-chat-app/chat"

func GetRooms() ([]chat.Room, error) {
	rooms := []chat.Room{}
	rows, err := db.Query("SELECT id, name from chatrooms")
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
