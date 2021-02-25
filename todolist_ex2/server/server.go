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

func ListenAndServe(cfg Config) error {
	err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		return err
	}

	return http.ListenAndServe(cfg.Address, loggingMiddleware(api.TodoListAPI()))

}
