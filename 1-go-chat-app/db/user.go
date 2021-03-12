package db

import (
	"github.com/jeonjonghyeok/go-run/1-go-chat-app/chat"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user chat.User) (int, error) {
	var id int
	passwd_hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return id, err
	}
	db.QueryRow("INSERT INTO users (username, password_hash) VALUES ($1, $2) RETURNING id", user.Username, passwd_hash).Scan(&id)
	return id, nil

}
