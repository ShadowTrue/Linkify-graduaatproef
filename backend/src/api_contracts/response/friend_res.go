package response

import (
	"time"

	"github.com/google/uuid"
)


type FriendRes struct {
	id uuid.UUID
	profilePicture string
	firstName string
	lastName string
	fullName string
	birthDate time.Time
}