package db

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jeonjonghyeok/go-run/1-go-chat-app/chat"
)

type ChatroomSubscription struct {
	Sub subscription
	C   chan chat.Message
}

func sendExitingMessage(chatid int, c chan chat.Message, limit int) error {
	rows, err := db.Query(`SELECT m.sender_id, u.username, m.text, m.sentOn
	FROM messages m
	JOIN users u ON m.chatroom_id = u.username
	WHERE m.chatroom_id = $1
	ORDER BY m.sent_on DESC 
	LIMIT $2)
	`, chatid, limit)
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

func NewChatroomSubscription(chatid int) (ChatroomSubscription, error) {
	log.Println("NewChatroomSubscription start")
	c := make(chan chat.Message, 128)
	err := sendExitingMessage(chatid, c, 100)
	if err != nil {
		return ChatroomSubscription{}, err
	}
	chatroomSubscription := ChatroomSubscription{
		Sub: subscribe(fmt.Sprintf("new_message_%d", chatid)),
		C:   c,
	}
	go func() {
		defer close(c)
		for m := range chatroomSubscription.Sub.c {
			var msg chat.Message
			if err := json.Unmarshal([]byte(m), &msg); err != nil {
				log.Println("unmarshalled error")
				continue
			}
			c <- msg
		}

	}()
	return chatroomSubscription, nil

}

func (c *ChatroomSubscription) Close() {
	c.Sub.close()
}

func SendMessage(userid int, chatid int, text string) error {
	_, err := db.Exec(`INSERT INTO messages (userid, chatid, text) VALUES ($1, $2)`, userid, chatid, text)
	return err
}
