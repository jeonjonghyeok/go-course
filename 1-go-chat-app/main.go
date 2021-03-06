package main

import (
	"net/http"

	"github.com/jeonjonghyeok/go-run/1-go-chat-app/ws"
)

func main() {
	http.ListenAndServe(":8080", ws.Handler())
}
