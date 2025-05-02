package chat_repo

import (
	"backend/src/models"

	"github.com/google/uuid"
)

type ChatRepo interface{
	GetChatById(chatId uuid.UUID) *models.Chat
	GetAllChatByUserId(userId uuid.UUID) *[]models.Chat
	CreateChat(chat models.Chat)
	UpdateChat(chatId models.Chat,chat models.Chat)
	DeleteChat(chatId models.Chat)
}