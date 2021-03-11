package chat

import "time"

type Message struct {
	Text   string    `json:"text"`
	Sender string    `json:"sender"`
	SentOn time.Time `json:"sentOn`
}
