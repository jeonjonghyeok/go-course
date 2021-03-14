package ws

import (
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/jeonjonghyeok/go-run/1-go-chat-app/chat"
	"github.com/jeonjonghyeok/go-run/1-go-chat-app/db"
)

const (
	readTimeout    = time.Second * 10
	writeTimeout   = time.Second * 10
	maxMessageSize = 512
	pingPeriod     = time.Second * 5
)

type Conn struct {
	wg     sync.WaitGroup
	ws     *websocket.Conn
	userid int
	chatid int
	sub    db.ChatroomSubscription
}

func newConn(ws *websocket.Conn, userid int, chatid int) *Conn {
	log.Println("create conn")
	return &Conn{
		ws:     ws,
		userid: userid,
		chatid: chatid,
	}
}

func (c *Conn) run() error {
	log.Println("run start")
	sub, err := db.NewChatroomSubscription(c.chatid)
	if err != nil {
		return err
	}
	c.sub = sub
	c.wg.Add(2)

	go c.readPump()
	go c.writePump()

	c.wg.Wait()
	return nil
}

//channel에서 데이터 읽는 go루틴
func (c *Conn) readPump() {
	defer c.wg.Done()
	defer c.sub.Close()

	c.ws.SetReadDeadline(time.Now().Add(readTimeout))
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetPongHandler(func(string) error {
		c.ws.SetReadDeadline(time.Now().Add(readTimeout))
		log.Println("get pong")
		return nil
	})

	for {
		var msg chat.Message
		err := c.ws.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			return
		}
		db.SendMessage(c.userid, c.chatid, msg.Text)
	}

}

func (c *Conn) writePump() {
	defer c.wg.Done()
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()

	for {
		select {
		case msg, more := <-c.sub.C:
			if !more {
				return
			}
			c.ws.SetWriteDeadline(time.Now().Add(writeTimeout))
			if err := c.ws.WriteJSON(msg); err != nil {
				log.Println("write JSON error")
				return
			}
		case <-ticker.C:
			log.Println("send ping")
			c.ws.WriteControl(
				websocket.PingMessage, nil, time.Now().Add(writeTimeout))
		}

	}
}
