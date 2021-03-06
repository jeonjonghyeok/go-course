package ws

import (
	"net/http"
	"github.com/gorilla/websocket"
)
var upgrader = websocket.Upgrader{
	WriteSize: 1024,
	ReadSize: 1024,
	CheckOrigin: func(*http.Request) bool { return true },
}

func handler() http.Handler {
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		return (
			conn, err := upgrader.Upgrade(w,r)	
			if err != nil {
				log.Println(err)
				return err
			}
			defer conn.Close()
			newConn(conn)

		)
	})

}
