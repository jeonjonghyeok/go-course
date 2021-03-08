package chat

import "time"

type Message struct {
	text   string    `json:"text"`
	sender string    `json:"sender"`
	sentOn time.Time `json:"sentOn`
}
