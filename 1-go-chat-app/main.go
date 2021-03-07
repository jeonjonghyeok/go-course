package main

import (
	"log"
	"net/http"

	"github.com/jeonjonghyeok/go-run/1-go-chat-app/ws"
)

func main() {
	if err := http.ListenAndServe(":5000", ws.Handler()); err != nil {
		log.Fatal(err)
	}
}
