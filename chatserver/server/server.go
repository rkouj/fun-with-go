package server

import (
	"rkouj/fun-with-go/chatserver/user"
	"rkouj/fun-with-go/chatserver/util"
)

type Server interface {
	ProcessRequests()
	GetMessages(u user.User) []string
}

type server struct {
	messages map[int][]string
	userReadChans []chan string
	userWriteChans []chan string
}

func NewServer(userReadChans []chan string, userWriteChans []chan string) Server {
	return &server{
		messages: make(map[int][]string),
		userReadChans: userReadChans,
		userWriteChans: userWriteChans,
	}
}

func (s *server) ProcessRequests() {
	for {
		select {
		case msg:= <- s.userWriteChans[0]:
			s.processMessage(msg)
		case msg:= <- s.userWriteChans[1]:
			s.processMessage(msg)
		}
	}
}

func (s *server) GetMessages(u user.User) []string {
	return s.messages[u.GetId()]
}

func (s *server) processMessage(msg string) {
	_, msgRecepientId, msgContent := util.DecipherMessage(msg)
	s.userReadChans[msgRecepientId] <- msg
	s.messages[msgRecepientId] = append(s.messages[msgRecepientId], msgContent)
}