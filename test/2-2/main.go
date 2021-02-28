package main

import (
	"github.com/gorilla/mux"
)

var upgrader = websocket.Upgrader{
	ReaderBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
	http.ListenAndServe(":8081", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request){
			conn, err := upgrader.Upgrade(w,r,nil)
			if err != nil {
				log.Fatalln(err)
			}

			typ, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("error reading",err)
			}
			if err := conn.WriteMessage(typ, msg); err != nil {
				log.Println("error reading",err)
			}
		}
	))
}
