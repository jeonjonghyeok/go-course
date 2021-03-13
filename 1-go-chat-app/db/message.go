package db

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jeonjonghyeok/go-run/1-go-chat-app/chat"
)

type ChatroomSubscription struct {
	sub subscription
	c   chan chat.Message
}

func NewChatroomSubscription(chatid int) (ChatroomSubscription, error) {
	c := make(chan chat.Message, 128)
	err := sendExitingMessage(chatid, c)
	if err != nil {
		return ChatroomSubscription{}, err
	}
	chatroomSubscription := ChatroomSubscription{
		sub: subscribe(fmt.Sprintf("new_message_%d", chatid)),
		c:   c,
	}
	go func() {
		var msg chat.Message
		for m := range chatroomSubscription.sub.c {
			if err := json.Unmarshal([]byte(m), &msg); err != nil {
				log.Println("unmarshalled error")
				continue
			}
			c <- msg
		}

	}()
	return chatroomSubscription, nil

}

func sendExitingMessage(chatid int, c chan chat.Message) error {
	rows, err := db.Query(`SELECT m.sender_id, u.username, m.text, m.sentOn
	FROM messages m
	JOIN users u ON m.chatroom_id = u.username
	WHERE m.chatroom_id = $1 
	`, chatid)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var m chat.Message
		err := rows.Scan(&m.Sender, &m.Text, &m.SentOn)
		if err != nil {
			return err
		}
		c <- m
	}
	return nil
}
