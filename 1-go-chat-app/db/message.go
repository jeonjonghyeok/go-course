package db

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jeonjonghyeok/go-run/1-go-chat-app/chat"
)

type ChatroomSubscription struct {
	sub subscription
	C   <-chan chat.Message
}

func sendExitingMessage(chatid int, c chan<- chat.Message,
	limit int) error {
	rows, err := db.Query(`WITH msgs AS (
		SELECT m.sender_id, u.username, m.text, m.sent_on
			FROM messages m
			JOIN users u ON m.sender_id = u.id
			WHERE m.chatroom_id = $1
			ORDER BY m.sent_on DESC 
			LIMIT $2)
		SELECT * FROM msgs ORDER BY sent_on ASC`, chatid, limit)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var m chat.Message
		err := rows.Scan(&m.SenderID, &m.Sender, &m.Text, &m.SentOn)
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
		sub: subscribe(fmt.Sprintf("new_message_%d", chatid)),
		C:   c,
	}
	go func() {
		defer close(c)
		for m := range chatroomSubscription.sub.c {
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
	c.sub.close()
}

func SendMessage(userid int, chatid int, text string) error {
	log.Println("Send Message start, userid=", userid, "chatid=", chatid)
	_, err := db.Exec(`INSERT INTO messages (sender_id, chatroom_id, text) 
	VALUES ($1, $2, $3)`, userid, chatid, text)
	return err
}
