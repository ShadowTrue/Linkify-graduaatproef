package response

import (
	"time"

	"github.com/google/uuid"
)

type MessageRes struct {
	Id uuid.UUID
	ChatId uuid.UUID
	SenderId uuid.UUID
	Content string
	TimeStamp time.Time
}
