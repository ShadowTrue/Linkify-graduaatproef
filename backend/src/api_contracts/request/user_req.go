package request

import (
	"time"

	"github.com/google/uuid"
)

type UserReq struct {
	Id             uuid.UUID
	ProfilePicture string
	Email          string
	Username       string
	FirstName      string
	LastName       string
	Birthday       time.Time
	Country        string
	CreatedOn      time.Time
}