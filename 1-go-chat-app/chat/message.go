package chat

import "time"

type Message struct {
	Text     string    `json:"text"`
	Sender   string    `json:"sender"`
	SenderID int       `json:"senderID"`
	SentOn   time.Time `json:"sentOn`
}

type Room struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
