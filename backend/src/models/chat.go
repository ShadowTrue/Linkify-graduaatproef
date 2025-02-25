package models

import (
	"errors"
	"github.com/google/uuid"
)

type chat struct {
	id           uuid.UUID
	participants []uuid.UUID
	admins       []uuid.UUID
}

func NewChat(id uuid.UUID, participants, admins []uuid.UUID) (*chat, []error) {

	if id == uuid.Nil {
		id = uuid.New()
	}
	newChat := &chat{id: id, participants: participants, admins: admins}

	ers := validateChat(newChat)
	if ers != nil {
		return nil, ers
	}
	return newChat, nil
}

func validateChat(c *chat) []error {
	var errs []error
	if c.participants == nil {
		errs = append(errs, errors.New("participants cannot be nil"))
	}
	if c.admins == nil {
		errs = append(errs, errors.New("admins cannot be nil"))
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}
