package ws

import "sync"

var conns = make(map[int]chan message)
var connsMux sync.RWMutex
var connCounter int

func addConn(c chan message) int {
	connsMux.Lock()
	defer connsMux.Unlock()

	id := connCounter
	conns[id] = c
	connCounter++
	return id
}

func deleteConn(id int) {
	connsMux.Lock()
	defer connsMux.Unlock()

	delete(conns, id)
}

func send(msg message) {
	connsMux.RLock()
	defer connsMux.RUnlock()

	for _, v := range conns {
		v <- msg
	}
}
