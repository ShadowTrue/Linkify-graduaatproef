package models

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type message struct {
	id        uuid.UUID
	chatId    uuid.UUID
	senderId  uuid.UUID
	content   string
	timeStamp time.Time
}

func NewMessage(id, chatId, senderId uuid.UUID, content string, timeStamp time.Time) (*message, []error) {

	newMessage := &message{id, chatId, senderId, content, timeStamp}
	errs := validateMessageParams(newMessage)

	if errs != nil {
		return nil, errs
	}

	return newMessage, nil
}

// Validates params and if errors occurs return slice with errors
func validateMessageParams(msg *message) []error {
	var errs []error

	if msg.id == uuid.Nil {
		errs = append(errs, errors.New("id is required"))
	}
	if msg.chatId == uuid.Nil {
		errs = append(errs, errors.New("chatId is required"))
	}
	if msg.senderId == uuid.Nil {
		errs = append(errs, errors.New("senderId is required"))
	}
	if strings.TrimSpace(msg.content) == "" {
		errs = append(errs, errors.New("content is required"))
	}
	if msg.timeStamp.IsZero() {
		errs = append(errs, errors.New("timeStamp is required"))
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}
