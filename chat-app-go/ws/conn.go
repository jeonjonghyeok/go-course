package ws

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type newconn struct {
	wsConn websocket.Conn
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
	defer c.wg.Done()
	go c.readPump()
	go c.writePump()
	c.wg.Wait()
	c.wsConn.Close()
}
func (c *newconn) readPump() {
	for {
		typ, msg, err := c.wsConn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		if typ != websocket.TextMessage {
			continue
		}
		log.Println("readPump msg= ", msg)
		c.ch <- msg
	}
}
func (c *newconn) writePump() {
	for {
		select {
		case msg, more := <-c.ch:
			if !more {
				return
			}
			log.Println("writePump msg= ", msg)
			err := c.wsConn.WriteMessage(websocket.TextMessage, msg)
			if err {
				log.Println(err)
				return
			}
		}

	}
}
