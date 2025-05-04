package message_repo

import (
	"backend/src/models"

	"github.com/google/uuid"
)

type IMessageRepo interface {
	GetMessageById(messageId uuid.UUID,page int,limit int) (*models.Message,error)
	GetMessagesByChat(chatId uuid.UUID) (*[]models.Message,error)
	CreateMessage(message models.Message) error
	UpdateMessage(message models.Message) error
	DeleteMesage(msgId uuid.UUID) error
}