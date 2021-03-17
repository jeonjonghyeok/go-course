package db

import (
	"database/sql"
	"errors"

	"github.com/jeonjonghyeok/go-run/1-go-chat-app/chat"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user chat.User) (id int, err error) {
	passwd_hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	err = db.QueryRow(`INSERT INTO users (username, password_hash) VALUES ($1, $2) RETURNING id`, user.Username, string(passwd_hash)).Scan(&id)
	return

}

var ErrUnauthorized = errors.New("db: unauthorized")

func FindUser(username string, password string) (id int, err error) {
	var password_hash []byte
	err = db.QueryRow(`SELECT id,password_hash FROM users WHERE username=$1`, username).Scan(&id, &password_hash)
	if err == sql.ErrNoRows {
		return 0, ErrUnauthorized
	}
	if err != nil {
		return 0, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(password)); err != nil {
		return 0, ErrUnauthorized
	}
	return
}
