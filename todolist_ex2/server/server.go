package server

import "net/http"

type Config struct {
	Address     string
	DatabaseURL string
}

func ListenAndServe(cfg Config) {
	db.Connect(cfg.DatabaseURL)
	http.ListenAndServe(cfg.Address, loggingMiddleware(api.TodoListAPI()))
}
