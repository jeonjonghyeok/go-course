package api

import (
	"net/http"

	"github.com/jeonjonghyeok/go-run/1-go-chat-app/chat"
	"github.com/jeonjonghyeok/go-run/1-go-chat-app/db"
)

func signup(w http.ResponseWriter, r *http.Request) {
	req := chat.User{}
	parseJSON(r.Body, &req)

	db.CreateUser(req)
}

func getRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := db.GetRooms()
	must(err)
	writeJSON(w, rooms)
}
