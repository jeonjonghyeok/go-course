package main

import (
	"github.com/jeonjonghyeok/todolist_ex2/server"
)

func main() {

	server.ListenAndServe(
		server.Config{
			Address:     ":5000",
			DatabaseURL: "postgres://postgres:tododbpasswd123@tododb.learningspoons.danslee.com:5432/postgres?sslmode=require",
		})
}
