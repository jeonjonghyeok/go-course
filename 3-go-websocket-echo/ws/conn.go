package ws

import (
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeTimeout   = 10 * time.Second
	readTimeout    = 10 * time.Second
	pingPeriod     = 5 * time.Second
	maxMessageSize = 512
)

type conn struct {
	wsConn *websocket.Conn
	send   chan []byte
	wg     sync.WaitGroup
}

func newConn(wsConn *websocket.Conn) *conn {
	return &conn{
		wsConn: wsConn,
		send:   make(chan []byte),
	}
}

func (c *conn) run() {
	c.wg.Add(2)
	go c.readPump()
	go c.writePump()
	c.wg.Wait()
	c.wsConn.Close()
}

func (c *conn) readPump() {
	defer c.wg.Done()

	c.wsConn.SetReadLimit(maxMessageSize)
	c.wsConn.SetReadDeadline(time.Now().Add(readTimeout))
	c.wsConn.SetPongHandler(func(string) error {
		log.Println("got pong")
		c.wsConn.SetReadDeadline(time.Now().Add(readTimeout))
		return nil
	})

	for {
		typ, msg, err := c.wsConn.ReadMessage()
		if err != nil {
			log.Println("err reading:", err)
			close(c.send)
			return
		}

		if typ != websocket.TextMessage {
			continue
		}

		c.send <- msg
	}
}

func (c *conn) writePump() {
	defer c.wg.Done()

	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()

	for {
		select {
		case s, more := <-c.send:
			if !more {
				return
			}

			log.Println("sending:", string(s))
			c.wsConn.SetWriteDeadline(time.Now().Add(writeTimeout))
			if err := c.wsConn.WriteMessage(websocket.TextMessage, s); err != nil {
				log.Println("err writing:", err)
				return
			}

		case <-ticker.C:
			log.Println("sent ping")
			c.wsConn.WriteControl(websocket.PingMessage, nil, time.Now().Add(writeTimeout))
		}
	}
}
