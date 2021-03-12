package db

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

var db *sql.DB
var listener *pq.Listener

type subscription struct {
	c chan string
}

var subscriptions map[string][]subscription

func Connect(url string) error {
	c, err := sql.Open("postgres", url)
	if err != nil {
		return err
	}
	db = c

	listener = pq.NewListener(url, time.Second*10, time.Minute, func(ev pq.ListenerEventType, err error) {
		if err != nil {
			panic(err)
		}
	})
	if subscriptions == nil {
		subscriptions = make(map[string][]subscription)
	}

	go func() {
		for n := range listener.NotificationChannel() {
			if channels, ok := subscriptions[n.Channel]; ok {
				for _, c := range channels {
					c.c <- n.Extra
				}
			}
		}
	}()

	return nil
}
