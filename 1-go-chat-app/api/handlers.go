package api

import (
	"log"
	"net/http"

	"github.com/jeonjonghyeok/go-run/1-go-chat-app/chat"
	"github.com/jeonjonghyeok/go-run/1-go-chat-app/db"
	"github.com/jeonjonghyeok/go-run/1-go-chat-app/token"
	"github.com/jeonjonghyeok/go-run/1-go-chat-app/ws"
)

func signup(w http.ResponseWriter, r *http.Request) {
	var req chat.User
	parseJSON(r.Body, &req)
	id, err := db.CreateUser(req)
	must(err)
	token, err := token.New(id)
	must(err)
	writeJSON(w, struct {
		Token string `json:"token"`
	}{token})
}

func signin(w http.ResponseWriter, r *http.Request) {
	req := chat.User{}
	parseJSON(r.Body, &req)
	id, err := db.FindUser(req.Username, req.Password)
	must(err)

	token, err := token.New(id)
	must(err)
	writeJSON(w, struct {
		Token string `json:"token"`
	}{token})
}

func createRoom(w http.ResponseWriter, r *http.Request) {
	log.Println("create Room")
	var req struct {
		Name string `json:"name"`
	}
	parseJSON(r.Body, &req)

	id, err := db.CreateRoom(req.Name)
	if err != nil {
		must(err)
	}
	writeJSON(w, struct {
		ID int `json:"id"`
	}{id})
}

func getRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := db.GetRooms()
	must(err)
	writeJSON(w, rooms)
}

func connectToRoom(w http.ResponseWriter, r *http.Request) {
	id, err := userID(r)
	if err != nil {
		must(err)
	}
	chatid := parseIntParam(r, "id")

	ws.ChatHandler(id, chatid).ServeHTTP(w, r)

}
