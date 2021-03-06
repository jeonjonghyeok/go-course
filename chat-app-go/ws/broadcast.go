package ws

import (
	"github.com/gorilla/mux"
)

var conns := make(map[int]chan []byte)
var chanCount int

func addConn(c chan []byte) {

	conns[chanCount++] = c
	
}