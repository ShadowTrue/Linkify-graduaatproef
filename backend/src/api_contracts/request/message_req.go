package request

import (
	"time"

	"github.com/google/uuid"
)

type MessageReq struct{
	chatId uuid.UUID
	sender uuid.UUID
	content string
	timeStamp time.Time
}