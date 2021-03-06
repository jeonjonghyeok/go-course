package ws

import (
	"sync"
)

var conns = make(map[int]chan []byte)
var connMux sync.RWMutex
var chanCount int

func addConn(c chan []byte) {
	connMux.Lock()
	defer connMux.Unlock()

	conns[chanCount] = c
	chanCount++
}

func send(msg []byte) {
	connMux.RLock()
	defer connMux.RUnlock()

	for _, c := range conns {
		c <- msg
	}
}
