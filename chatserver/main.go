package main

import (
	"rkouj/fun-with-go/chatserver/user"
	"rkouj/fun-with-go/chatserver/server"
	"time"
)

const (
	numUsers = 2
	bufferSize = 100
	sleepDuration = 1 * time.Millisecond
)

func main() {

	var (
		users = make([]user.User, numUsers)
		readChans = make([]chan string, numUsers)
		writeChans = make([]chan string, numUsers)
	)

	for i:=0; i<numUsers; i++ {
		readChans[i] = make(chan string, bufferSize)
		writeChans[i] = make(chan string, bufferSize)
	}

	serverObject := server.NewServer(readChans, writeChans)
	go serverObject.ProcessRequests()

	for u := range users {
		users[u] = user.NewUser(u, readChans[u], writeChans[u])
		go users[u].StartChatting()
	}

	readChans[0] <- "0:1:TestMessage"

	time.Sleep(sleepDuration)

	for i:=1; i<numUsers; i++ {
		close(readChans[i])
		close(writeChans[i])
	}
}
