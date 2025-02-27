package request

import (
	"time"

	"github.com/google/uuid"
)

type MessageReq struct{
	ChatId uuid.UUID
	Sender uuid.UUID
	Content string
	TimeStamp time.Time
}