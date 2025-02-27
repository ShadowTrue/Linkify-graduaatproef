package response

import (
	"time"

	"github.com/google/uuid"
)

type MessageRes struct {
	id uuid.UUID
	chatId uuid.UUID
	senderId uuid.UUID
	content string
	timeStamp time.Time
}
