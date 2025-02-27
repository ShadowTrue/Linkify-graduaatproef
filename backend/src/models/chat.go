package models

import (
	"errors"

	"github.com/google/uuid"
)

type Chat struct {
	Id           uuid.UUID
	Participants []uuid.UUID
	Admins       []uuid.UUID
}

func NewChat(id uuid.UUID, participants, admins []uuid.UUID) (*Chat, []error) {

	if id == uuid.Nil {
		id = uuid.New()
	}
	newChat := &Chat{Id: id, Participants: participants, Admins: admins}

	ers := validateChat(newChat)
	if ers != nil {
		return nil, ers
	}
	return newChat, nil
}

func validateChat(c *Chat) []error {
	var errs []error
	if c.Participants == nil {
		errs = append(errs, errors.New("participants cannot be nil"))
	}
	if c.Admins == nil {
		errs = append(errs, errors.New("admins cannot be nil"))
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}
