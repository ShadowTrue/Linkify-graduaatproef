package response

import (
	"time"

	"github.com/google/uuid"
)

type PostRes struct {
	Id uuid.UUID
	Sender uuid.UUID
	Content string
	Image string
	TimeStamp time.Time
}