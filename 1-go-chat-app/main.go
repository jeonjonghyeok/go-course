package main

import (
	"github.com/jeonjonghyeok/go-run/1-go-chat-app/server"
)

func main() {
	server.ListenAndServe(server.Config{
		Address: ":5000",
		Url:     "postgres://postgres:chatdbpasswd123@chatdb.learningspoons.danslee.com:5432/postgres",
	})
}
