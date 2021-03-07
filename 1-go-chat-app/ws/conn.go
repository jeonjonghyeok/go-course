package ws

import (
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	readTimeout    = 10 * time.Second
	writeTimeout   = 10 * time.Second
	maxMessageSize = 512
	pingPeriod     = 5 * time.Second
)

type conn struct {
	wsConn *websocket.Conn
	send   chan []byte
	wg     sync.WaitGroup
}

func newConn(c *websocket.Conn) *conn {
	return &conn{
		wsConn: c,
		send:   make(chan []byte),
	}
}

func (c *conn) run() {
	c.wg.Add(2)
	go readPump(c)
	go writePump(c)

	c.wg.Wait()
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
		typ, msg, err := c.wsConn.ReadMessage()
		if err != nil {
			close(c.send)
			log.Println(err)
			return
		}
		if typ != websocket.TextMessage {
			continue
		}
		c.send <- msg

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
			log.Println("msg= ", string(msg))
			c.wsConn.SetWriteDeadline(time.Now().Add(writeTimeout))
			if err := c.wsConn.WriteMessage(websocket.TextMessage, msg); err != nil {
				log.Println(err)
				return
			}

		case <-ticker.C:
			log.Println("send ping")
			c.wsConn.WriteControl(websocket.PingMessage, nil, time.Now().Add(writeTimeout))
		}
	}

}
