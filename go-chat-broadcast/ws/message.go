package ws

import "time"

type message struct {
	Sender string    `json:"sender"`
	Text   string    `json:"text"`
	SentOn time.Time `json:"sentOn"`
}
