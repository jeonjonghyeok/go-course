package main

import (
	"log"
	"net/http"

	"github.com/learningspoons-go/websocket-chat/ws"
)

func main() {
	if err := http.ListenAndServe(":5000", ws.Handler()); err != nil {
		log.Fatalln(err)
	}
}
