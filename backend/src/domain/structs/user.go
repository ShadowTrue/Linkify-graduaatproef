package user

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

//Pointers are set so that params can be nil
//Some fields are not required for example when registering
//Id is only generated after succesfully registering same for createdOn
type user struct {
	id *uuid.UUID
	email string 
	username string
	firstName string
	lastName string
	birthday time.Time
	country *string
	createdOn *time.Time
}
// Constructor for User
func NewUser(id *uuid.UUID, email string,
	 username string,firstName string,
	lastName string, birthday time.Time,
	 country *string, createdOn *time.Time) (*user, []error){


	//validates user parameters
	errs := validateParams(email,username,firstName,lastName,birthday)
	
	if(errs != nil){
		return nil, errs
	}

	if(id == nil){
		newId := uuid.New()
		id = &newId
	}

	if(createdOn == nil || createdOn.IsZero()){
		now := time.Now().UTC()
		createdOn = &now
	}

	return &user{id,email,username,firstName,lastName,birthday,country,createdOn},nil
}

//Validation
func validateParams(email ,username ,firstName ,lastName string, birthday time.Time) []error{
	var errs []error 
	if(strings.TrimSpace(email) == ""){
		errs = append(errs,errors.New("Email cannot be empty"))
	}

	if(strings.TrimSpace(username) == ""){
		errs = append(errs,errors.New("Username cannot be empty"))
	}

	if(strings.TrimSpace(firstName) == ""){
		errs = append(errs,errors.New("First name cannot be empty"))
	}

	if(strings.TrimSpace(lastName) == ""){
		errs = append(errs,errors.New("Last name cannot be empty"))
	}

	if(birthday.IsZero()){
		errs = append(errs,errors.New("Birthday need to be set"))
	}
	
	if(len(errs) == 0){
		return nil
	}

	return errs
}



