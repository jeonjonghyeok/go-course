package server

import (
	"net/http"

	"github.com/jeonjonghyeok/todolist_test/api"
	"github.com/jeonjonghyeok/todolist_test/db"
)

type Config struct {
	Address     string
	DatabaseURL string
}

func ListenAndServe(cfg Config) error {
	if err := db.Connect(cfg.DatabaseURL); err != nil {
		return err
	}
	return http.ListenAndServe(cfg.Address, LoggingHandler(api.TodoListAPI()))
}
