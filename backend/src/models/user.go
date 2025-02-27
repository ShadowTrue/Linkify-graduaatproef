package models

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type User struct {
	id             uuid.UUID
	profilePicture string
	email          string
	username       string
	firstName      string
	lastName       string
	birthday       time.Time
	country        string
	createdOn      time.Time
}

// Constructor for User
func NewUser(id uuid.UUID, profilePicture, email, username, firstName, lastName string, birthday time.Time, country string, createdOn time.Time) (*User, []error) {

	//assigns new UUID
	if id == uuid.Nil {
		id = uuid.New()
	}
	//set creation date
	if createdOn.IsZero() {
		createdOn = time.Now().UTC()
	}
	newUser := &User{id, profilePicture, email, username, firstName, lastName, birthday, country, createdOn}
	errs := validateUserParams(newUser)

	if errs != nil {
		return nil, errs
	}
	return newUser, nil
}

// Validation
func validateUserParams(u *User) []error {
	var errs []error
	if strings.TrimSpace(u.email) == "" {
		errs = append(errs, errors.New("email cannot be empty"))
	}

	if strings.TrimSpace(u.username) == "" {
		errs = append(errs, errors.New("username cannot be empty"))

	}

	if strings.TrimSpace(u.firstName) == "" {
		errs = append(errs, errors.New("first name cannot be empty"))
	}

	if strings.TrimSpace(u.lastName) == "" {
		errs = append(errs, errors.New("last name cannot be empty"))
	}
	if u.birthday.IsZero() {
		errs = append(errs, errors.New("birthday need to be set"))
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}
