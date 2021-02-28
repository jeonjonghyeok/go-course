package main

import (
	"log"

	"github.com/jeonjonghyeok/todolist_test/server"
)

func main() {
	if err := server.ListenAndServe(server.Config{
		Address:     ":8086",
		DatabaseURL: "postgres://postgres:tododbpasswd123@tododb.learningspoons.danslee.com:5432/postgres?sslmode=require",
	}); err != nil {
		log.Fatalln(err)
	}
}
