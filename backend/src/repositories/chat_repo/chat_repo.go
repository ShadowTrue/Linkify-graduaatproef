package chat_repo

import (
	"backend/src/models"

	"github.com/google/uuid"
)

type ChatRepo interface{
	GetChatById(chatId uuid.UUID) (*models.Chat,error)
	GetAllChatByUserId(userId uuid.UUID) (*[]models.Chat,error)
	CreateChat(chat models.Chat) error
	UpdateChat(chatId models.Chat,chat models.Chat) error
	DeleteChat(chatId models.Chat) error
}