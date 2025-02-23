package models

import (
	"errors"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type user struct {
	id             uuid.UUID
	profilePicture string
	email          string `validate:"required,email"`
	username       string `validate:"required"`
	firstName      string `validate:"required"`
	lastName       string
	hashedPassword string    `validate:"required"`
	birthday       time.Time `validate:"required,Datetime"`
	country        string
	createdOn      time.Time
}

// Constructor for User
func NewUser(id uuid.UUID, profilePicture, email, hashedPassword, username, firstName, lastName string, birthday time.Time, country *string, createdOn *time.Time) (*user, []error) {

	//validates user parameters
	errs := validateUserParams(email, username, firstName, hashedPassword, lastName, birthday)

	//return slice of errors
	if errs != nil {
		return nil, errs
	}

	//assigns new UUID
	if id == uuid.Nil {
		id = uuid.New()
	}
	//set creation date
	if createdOn == nil || createdOn.IsZero() {
		now := time.Now().UTC()
		createdOn = &now
	}

	return &user{id, profilePicture, email, username, firstName, lastName, hashedPassword, birthday, country, createdOn}, nil
}

// Validation
func validateUserParams(email, username, firstName, hashedPassword, lastName string, birthday time.Time) []error {
	var errs []error
	if strings.TrimSpace(email) == "" {
		errs = append(errs, errors.New("Email cannot be empty"))
	}

	if strings.TrimSpace(username) == "" {
		errs = append(errs, errors.New("Username cannot be empty"))

	}

	if strings.TrimSpace(firstName) == "" {
		errs = append(errs, errors.New("First name cannot be empty"))
	}

	if strings.TrimSpace(lastName) == "" {
		errs = append(errs, errors.New("Last name cannot be empty"))
	}
	if strings.TrimSpace(hashedPassword) == "" {
		errs = append(errs, errors.New("Password cannot be empty"))
	}
	if birthday.IsZero() {
		errs = append(errs, errors.New("Birthday need to be set"))
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}
