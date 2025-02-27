package response

import (
	"time"

	"github.com/google/uuid"
)

type UserRes struct {
	id uuid.UUID
	profilePicture string
	email string
	username string
	firstName string
	lastName string
	birthday time.Time
	country string
	createdOn time.Time

}