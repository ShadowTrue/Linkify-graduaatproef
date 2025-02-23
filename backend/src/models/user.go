package models

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type user struct {
	id             uuid.UUID
	profilePicture string
	email          string
	username       string
	firstName      string
	lastName       string
	hashedPassword string
	birthday       time.Time
	country        string
	createdOn      time.Time
}

// Constructor for User
func NewUser(id uuid.UUID, profilePicture, email, hashedPassword, username, firstName, lastName string, birthday time.Time, country string, createdOn time.Time) (*user, []error) {

	//assigns new UUID
	if id == uuid.Nil {
		id = uuid.New()
	}
	//set creation date
	if createdOn.IsZero() {
		createdOn = time.Now().UTC()
	}
	newUser := &user{id, profilePicture, email, username, firstName, lastName, hashedPassword, birthday, country, createdOn}
	errs := validateUserParams(newUser)

	if errs != nil {
		return nil, errs
	}
	return newUser, nil
}

// Validation
func validateUserParams(u *user) []error {
	var errs []error
	if strings.TrimSpace(u.email) == "" {
		errs = append(errs, errors.New("Email cannot be empty"))
	}

	if strings.TrimSpace(u.username) == "" {
		errs = append(errs, errors.New("Username cannot be empty"))

	}

	if strings.TrimSpace(u.firstName) == "" {
		errs = append(errs, errors.New("First name cannot be empty"))
	}

	if strings.TrimSpace(u.lastName) == "" {
		errs = append(errs, errors.New("Last name cannot be empty"))
	}
	if strings.TrimSpace(u.hashedPassword) == "" {
		errs = append(errs, errors.New("Password cannot be empty"))
	}
	if u.birthday.IsZero() {
		errs = append(errs, errors.New("Birthday need to be set"))
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}
