package ws

/*
import (
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/jeonjonghyeok/go-run/1-go-chat-app/chat"
)

const (
	readTimeout    = 10 * time.Second
	writeTimeout   = 10 * time.Second
	maxMessageSize = 512
	pingPeriod     = 5 * time.Second
)

type conn struct {
	wsConn *websocket.Conn
	send   chan chat.Message
	wg     sync.WaitGroup
}

func (c *conn) Send(msg chat.Message) {
	c.send <- msg
}

func newConn(c *websocket.Conn) *conn {
	return &conn{
		wsConn: c,
		send:   make(chan chat.Message),
	}
}

func (c *conn) run() {
	c.wg.Add(2)

	go readPump(c)
	go writePump(c)
	id := chat.GetRoom().AddParticipant(c)

	c.wg.Wait()
	chat.GetRoom().RemoveParticipant(id)
	c.wsConn.Close()
}

func readPump(c *conn) {
	defer c.wg.Done()

	c.wsConn.SetReadDeadline(time.Now().Add(readTimeout))
	c.wsConn.SetReadLimit(maxMessageSize)
	c.wsConn.SetPongHandler(func(string) error {
		log.Println("get pong")
		c.wsConn.SetReadDeadline(time.Now().Add(readTimeout))
		return nil
	})

	for {
		var msg chat.Message
		err := c.wsConn.ReadJSON(&msg)
		if err != nil {
			close(c.send)
			log.Println(err)
			return
		}
		log.Println(msg)

		chat.GetRoom().SendMessage(msg)
		//c.send <- msg

	}

}

func writePump(c *conn) {
	defer c.wg.Done()

	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()
	for {
		select {
		case msg, more := <-c.send:
			if !more {
				return
			}
			c.wsConn.SetWriteDeadline(time.Now().Add(writeTimeout))
			if err := c.wsConn.WriteJSON(msg); err != nil {
				log.Println(err)
				return
			}

		case <-ticker.C:
			log.Println("send ping")
			c.wsConn.WriteControl(websocket.PingMessage, nil, time.Now().Add(writeTimeout))
		}
	}

}
*/
