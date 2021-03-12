package api

import (
	"net/http"

	"github.com/jeonjonghyeok/go-run/1-go-chat-app/chat"
	"github.com/jeonjonghyeok/go-run/1-go-chat-app/db"
	"github.com/jeonjonghyeok/go-run/1-go-chat-app/token"
)

func signup(w http.ResponseWriter, r *http.Request) {
	req := chat.User{}
	parseJSON(r.Body, &req)

	id, err := db.CreateUser(req)
	must(err)
	token, err := token.New(id)
	must(err)
	writeJSON(w, token)
}

func signin(w http.ResponseWriter, r *http.Request) {

}

func getRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := db.GetRooms()
	must(err)
	writeJSON(w, rooms)
}
