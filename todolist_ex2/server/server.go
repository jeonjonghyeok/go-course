package server

import (
	"net/http"

	"github.com/jeonjonghyeok/todolist_ex2/api"
	"github.com/jeonjonghyeok/todolist_ex2/db"
)

type Config struct {
	Address     string
	DatabaseURL string
}

func ListenAndServe(cfg Config) {
	db.Connect(cfg.DatabaseURL)
	http.ListenAndServe(cfg.Address, loggingMiddleware(api.TodoListAPI()))
}
