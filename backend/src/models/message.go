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
	errs := validateMessageParams(id, chatId, senderId, content, timeStamp)

	if errs != nil {
		return nil, errs
	}

	return &message{id, chatId, senderId, content, timeStamp}, nil
}

// Validates params and if errors occurs return slice with errors
func validateMessageParams(id, chatId, senderId uuid.UUID, content string, timeStamp time.Time) []error {
	var errs []error

	if id == uuid.Nil {
		errs = append(errs, errors.New("id is required"))
	}
	if chatId == uuid.Nil {
		errs = append(errs, errors.New("chatId is required"))
	}
	if senderId == uuid.Nil {
		errs = append(errs, errors.New("senderId is required"))
	}
	if strings.TrimSpace(content) == "" {
		errs = append(errs, errors.New("content is required"))
	}
	if timeStamp.IsZero() {
		errs = append(errs, errors.New("timeStamp is required"))
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}
