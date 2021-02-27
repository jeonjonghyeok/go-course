package main

import (
	"log"

	"github.com/jeonjonghyeok/todolist_ex2/server"
)

func main() {

	err := server.ListenAndServe(
		server.Config{
			Address:     ":5000",
			DatabaseURL: "postgres://postgres:tododbpasswd123@tododb.learningspoons.danslee.com:5432/postgres?sslmode=require",
		})
	if err != nil {
		log.Fatalln(err)
	}
}
