package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	if err := http.ListenAndServe(":5000", handler()); err != nil {
		log.Fatalln(err)
	}
}

func handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()

		for {
			typ, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println("msg: ", string(msg))
			if err := conn.WriteMessage(typ, msg); err != nil {
				log.Println(err)
				return
			}
		}
	})
}
