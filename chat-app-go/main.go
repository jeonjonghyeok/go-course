package main

import (
	"log"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":5000", messageHandler()); err != nil {
		log.Fatalln(err)
	}
}
