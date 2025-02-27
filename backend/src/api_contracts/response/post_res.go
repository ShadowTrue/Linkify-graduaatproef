package response

import (
	"time"

	"github.com/google/uuid"
)

type PostRes struct {
	id uuid.UUID
	sender uuid.UUID
	content string
	image string
	timeStamp time.Time
}