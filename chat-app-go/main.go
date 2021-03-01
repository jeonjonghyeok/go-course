package main

import (
	"log"
	"net/http"

	"github.com/jeonjonghyeok/go-run/chat-app-go/ws"
)

func main() {
	if err := http.ListenAndServe(":5000", ws.MessageHandler()); err != nil {
		log.Fatalln(err)
	}
}
