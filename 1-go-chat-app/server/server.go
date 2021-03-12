package server

import (
	"log"
	"net/http"

	"github.com/jeonjonghyeok/go-run/1-go-chat-app/api"
	"github.com/jeonjonghyeok/go-run/1-go-chat-app/db"
)

type Config struct {
	Address string
	Url     string
}

func ListenAndServe(c Config) {
	db.Connect(c.Url)

	if err := http.ListenAndServe(":5000", api.ChatAPI()); err != nil {
		log.Fatal(err)
	}
}
