package models

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Friend struct {
	id             uuid.UUID
	profilePicture string
	firstName      string
	lastName       string
	fullName       string
	birthDate      time.Time
}

func NewFriend(id uuid.UUID, profilePicture, firstName, lastName, fullName string, birthday time.Time) (*Friend, []error) {

	if fullName == "" {
		fullName = firstName + lastName
	}
	newFriend := &Friend{id, profilePicture, firstName, lastName, fullName, birthday}
	errs := validateFriendParams(newFriend)

	if errs != nil {
		return nil, errs
	}
	return newFriend, nil
}

func validateFriendParams(f *Friend) []error {

	var errs []error
	if f.id == uuid.Nil {
		errs = append(errs, errors.New("id is required"))
	}
	if strings.TrimSpace(f.firstName) == "" {
		errs = append(errs, errors.New("firstName is required"))
	}
	if strings.TrimSpace(f.lastName) == "" {
		errs = append(errs, errors.New("lastName is required"))
	}
	if strings.TrimSpace(f.fullName) == "" {
		errs = append(errs, errors.New("fullName is required"))
	}
	if f.birthDate.IsZero() {
		errs = append(errs, errors.New("birthDate is required"))
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}
