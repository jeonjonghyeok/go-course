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

type newconn struct {
	wsConn *websocket.Conn
	ch     chan []byte
	wg     sync.WaitGroup
}

func newConn(conn *websocket.Conn) *newconn {
	return &newconn{
		wsConn: conn,
		ch:     make(chan []byte),
	}
}

func (c *newconn) run() {
	c.wg.Add(2)

	addConn(c.ch)

	go c.readPump()
	go c.writePump()
	c.wg.Wait()
	c.wsConn.Close()
}
func (c *newconn) readPump() {
	defer c.wg.Done()

	c.wsConn.SetReadLimit(maxMessageSize)
	c.wsConn.SetReadDeadline(time.Now().Add(readTimeout))
	c.wsConn.SetPongHandler(func(string) error {
		log.Println("get pong")
		c.wsConn.SetReadDeadline(time.Now().Add(readTimeout))
		return nil
	})

	for {
		typ, msg, err := c.wsConn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		if typ != websocket.TextMessage {
			continue
		}
		log.Println("readPump msg= ", string(msg))

		send(msg)
		//c.ch <- msg
	}
}
func (c *newconn) writePump() {
	defer c.wg.Done()

	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()

	for {
		select {
		case msg, more := <-c.ch:
			if !more {
				return
			}
			log.Println("writePump msg= ", string(msg))
			c.wsConn.SetWriteDeadline(time.Now().Add(writeTimeout))
			err := c.wsConn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println(err)
				return
			}
		case <-ticker.C:
			c.wsConn.WriteControl(websocket.PingMessage, nil, time.Now().Add(writeTimeout))
			log.Println("send ping")
		}

	}
}
