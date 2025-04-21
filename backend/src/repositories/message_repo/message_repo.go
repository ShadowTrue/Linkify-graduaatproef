package message_repo

import (
	"backend/src/models"

	"github.com/google/uuid"
)

type MessageRepo interface {
	GetMessageById(messageId uuid.UUID,page int,limit int) *models.Message
	GetMessagesByChat(chatId uuid.UUID) *[]models.Message
	CreateMessage(message models.Message)
	UpdateMessage(message models.Message)
	DeleteMesage(msgId uuid.UUID)
}