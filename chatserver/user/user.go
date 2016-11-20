package user

import (
	"rkouj/fun-with-go/chatserver/util"
	"fmt"
)

type User interface {
	StartChatting()
	GetId() int
}

type user struct {
	id int
	messagesSentCount int
	messagesReceivedCount int
	readChan chan string
	writeChan chan string
}

func NewUser(id int, readChan chan string, writeChan chan string) User {
	return &user{
		id: id,
		messagesSentCount: 0,
		messagesReceivedCount: 0,
		readChan: readChan,
		writeChan: writeChan,
	}
}

func (u *user) StartChatting() {
	for {
		select {
		case msg := <-u.readChan:
			u.messagesReceivedCount++
			fmt.Printf("User %d received message %d times\n", u.id, u.messagesReceivedCount)
			msgSenderId, messageRecepientId, messageContent := util.DecipherMessage(msg)
			u.writeChan <- fmt.Sprintf("%d%s%d%s%s", messageRecepientId, util.Delimiter, msgSenderId, util.Delimiter, u.generateReply(messageContent))
		}
	}
}

func (u *user) GetId() int {
	return u.id
}

func (u *user) generateReply(incomingMessage string) string {
	u.messagesSentCount++
	return fmt.Sprintf("User id %d is sending message number %d", u.id, u.messagesSentCount)
}