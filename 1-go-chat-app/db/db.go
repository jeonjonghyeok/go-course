package db

import (
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/lib/pq"
)

var db *sql.DB
var listener *pq.Listener

type subscription struct {
	name string
	c    chan string
}

var subscriptions map[string][]subscription
var subscriptionsMux sync.Mutex

func Connect(url string) error {
	c, err := sql.Open("postgres", url)
	if err != nil {
		return err
	}
	db = c
	subscriptions = make(map[string][]subscription)

	listener = pq.NewListener(url, time.Second*10, time.Minute, func(ev pq.ListenerEventType, err error) {
		if err != nil {
			panic(err)
		}
	})

	go func() {
		for n := range listener.NotificationChannel() {
			if channels, ok := subscriptions[n.Channel]; ok {
				for _, c := range channels {
					log.Println("notification")
					c.c <- n.Extra
				}
			}
		}
	}()

	return nil
}

func subscribe(name string) subscription {
	subscriptionsMux.Lock()
	defer subscriptionsMux.Unlock()

	if subscriptions[name] == nil {
		subscriptions[name] = []subscription{}
		err := listener.Listen(name)
		if err != nil {
			panic(err)
		}
	}
	c := subscription{
		name: name,
		c:    make(chan string, 256),
	}

	subscriptions[name] = append(subscriptions[name], c)
	return c
}

func (c *subscription) close() {
	subscriptionsMux.Lock()
	defer subscriptionsMux.Unlock()
	j := 0
	for _, subscriptionChannel := range subscriptions[c.name] {
		if subscriptionChannel.c != c.c {
			subscriptions[c.name][j] = subscriptionChannel
			j++
		}
	}
	subscriptions[c.name] = subscriptions[c.name][:j]
	close(c.c)

	if len(subscriptions[c.name]) == 0 {
		err := listener.Unlisten(c.name)
		if err != nil {
			panic(err)
		}
	}
	subscriptions[c.name] = nil
}
