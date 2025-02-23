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

	var errs []error
	if id == uuid.Nil {
		id = uuid.New()
	}
	if participants == nil {
		errs = append(errs, errors.New("participants cannot be nil"))
	}
	if admins == nil {
		errs = append(errs, errors.New("admins cannot be nil"))
	}
	if len(errs) > 0 {
		return nil, errs
	}
	return &chat{id: id, participants: participants, admins: admins}, nil
}
