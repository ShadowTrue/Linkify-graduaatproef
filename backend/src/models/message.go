package models

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id        uuid.UUID
	ChatId    uuid.UUID
	SenderId  uuid.UUID
	Content   string
	TimeStamp time.Time
}

func NewMessage(id, chatId, senderId uuid.UUID, content string, timeStamp time.Time) (*Message, []error) {

	if(id == uuid.Nil){
		id = uuid.New()
	}
	newMessage := &Message{id, chatId, senderId, content, timeStamp}
	errs := validateMessageParams(newMessage)

	if errs != nil {
		return nil, errs
	}

	return newMessage, nil
}

// Validates params and if errors occurs return slice with errors
func validateMessageParams(msg *Message) []error {
	var errs []error

	if msg.Id == uuid.Nil {
		errs = append(errs, errors.New("id is required"))
	}
	if msg.ChatId == uuid.Nil {
		errs = append(errs, errors.New("chatId is required"))
	}
	if msg.SenderId == uuid.Nil {
		errs = append(errs, errors.New("senderId is required"))
	}
	if strings.TrimSpace(msg.Content) == "" {
		errs = append(errs, errors.New("content is required"))
	}
	if msg.TimeStamp.IsZero() {
		errs = append(errs, errors.New("timeStamp is required"))
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}
