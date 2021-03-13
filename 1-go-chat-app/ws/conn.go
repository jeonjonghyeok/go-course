package ws

import (
	"sync"

	"github.com/gorilla/websocket"
	"github.com/jeonjonghyeok/go-run/1-go-chat-app/db"
)

type Conn struct {
	wg     sync.WaitGroup
	ws     *websocket.Conn
	userid int
	chatid int
	sub    db.ChatroomSubscription
}

func newConn(ws *websocket.Conn, userid int, chatid int) Conn {
	return Conn{
		ws:     ws,
		userid: userid,
		chatid: chatid,
	}
}

func (c Conn) run() {
	sub, err := db.NewChatroomSubscription(c.chatid)
	if err != nil {
		return
	}
	c.sub = sub
	c.wg.Add(2)

	go readPump(c)
	go writePump(c)

	c.wg.Wait()
}

//channel에서 데이터 읽는 go루틴
func readPump(c Conn) {
	defer c.wg.Done()

}

func writePump(c Conn) {
	defer c.wg.Done()
}
