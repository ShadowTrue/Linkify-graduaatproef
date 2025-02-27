package response

import (
	"time"

	"github.com/google/uuid"
)


type FriendRes struct {
	Id uuid.UUID
	ProfilePicture string
	FirstName string
	LastName string
	FullName string
	BirthDate time.Time
}