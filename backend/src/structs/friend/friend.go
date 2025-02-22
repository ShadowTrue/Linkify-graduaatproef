package friend

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type friend struct {
	id        uuid.UUID
	firstName string
	lastName  string
	fullName  string
	birthDate time.Time
}

func NewFriend(id uuid.UUID, firstName, lastName, fullName string, birthday time.Time) (*friend, []error) {

	if fullName == "" {
		fullName = firstName + lastName
	}

	errs := validateParams(id, firstName, lastName, fullName, birthday)

	if errs != nil {
		return nil, errs
	}
	return &friend{id, firstName, lastName, fullName, birthday}, nil
}

func validateParams(id uuid.UUID, firstName, lastName, fullName string, birthday time.Time) []error {

	var errs []error
	if id == uuid.Nil {
		errs = append(errs, errors.New("id is required"))
	}
	if firstName == "" {
		errs = append(errs, errors.New("firstName is required"))
	}
	if lastName == "" {
		errs = append(errs, errors.New("lastName is required"))
	}
	if fullName == "" {
		errs = append(errs, errors.New("fullName is required"))
	}
	if birthday.IsZero() {
		errs = append(errs, errors.New("birthDate is required"))
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}
