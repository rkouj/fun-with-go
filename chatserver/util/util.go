package util

import (
	"strings"
	"strconv"
	"fmt"
)

const (
	Delimiter = ":"
)

func DecipherMessage(message string) (int, int, string) {
	content := strings.Split(message, Delimiter)
	senderId, err1 := strconv.Atoi(content[0])
	receiverId, err2 := strconv.Atoi(content[1])
	if err1 != nil || err2 != nil {
		panic(fmt.Sprintf("Invalid message format %v %v %v", message, err1, err2))
	}

	return senderId, receiverId, content[1]
}