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
	if err := db.Connect(c.Url); err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(c.Address, api.ChatAPI()); err != nil {
		log.Fatal(err)
	}
}
