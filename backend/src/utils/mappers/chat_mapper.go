package mappers

import (
	"backend/src/api_contracts/request"
	"backend/src/api_contracts/response"
	"backend/src/models"
)

func MapChatToStruct(req *request.ChatReq) (models.Chat,[]error){
	chat,errs := models.NewChat(req.Id,req.Participants,req.Admins)
	return *chat,errs
}

func MapChatToRes(chat *models.Chat) response.ChatRes{
	res := response.ChatRes{Id: chat.Id,Participants: chat.Participants,Admins: chat.Admins}

	return res
}