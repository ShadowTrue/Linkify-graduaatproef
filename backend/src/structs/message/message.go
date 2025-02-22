package message

import (
	"time"

	"github.com/google/uuid"
)

//"errors"
//"strings"
//

type message struct {
	id        uuid.UUID
	chatId    uuid.UUID
	senderId  uuid.UUID
	content   string
	timeStamp time.Time
}

func NewMessage(id, chatId, senderId uuid.UUID, content string, timeStamp time.Time) (*message, []error) {

}

func validateParams(id, chatId, senderId uuid.UUID, content string, timeStamp time.Time) []error {

}
