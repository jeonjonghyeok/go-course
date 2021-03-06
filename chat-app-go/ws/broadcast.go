package ws

import (
	"sync"
)

var conns = make(map[int]chan []byte)
var connMux sync.RWMutex
var chanCount int

func addConn(c chan []byte) (id int) {
	connMux.Lock()
	defer connMux.Unlock()

	conns[chanCount] = c
	id = chanCount
	chanCount++
	return
}

func send(msg []byte) {
	connMux.RLock()
	defer connMux.RUnlock()

	for _, c := range conns {
		c <- msg
	}
}

func deleteConn(id int) {
	connMux.Lock()
	defer connMux.Unlock()

	delete(conns, id)

}
