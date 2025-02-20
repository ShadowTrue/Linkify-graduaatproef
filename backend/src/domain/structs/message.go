package structs

import (
	"time"

	"github.com/google/uuid"
)


type User struct {
	id uuid.UUID
	email string 
	username string
	firstName string
	lastName string
	birthday time.Time
	country string
	createdOn time.Time
}

func NewUser(id string, email string,
	 username string,firstName string,
	lastName string, birthday time.Time,
	 country string, createdOn time.Time) (*User, error){




}



