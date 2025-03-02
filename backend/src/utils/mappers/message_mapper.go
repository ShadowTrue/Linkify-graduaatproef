package mappers

import (
	"backend/src/api_contracts/request"
	"backend/src/api_contracts/response"
	"backend/src/models"
)

func MapMessageToStruct(req *request.MessageReq) (models.Message,[]error){
	message,errs := models.NewMessage(req.Id,req.ChatId,req.Sender,req.Content,req.TimeStamp)

	return *message,errs
}

func MapMessageToRes(msg *models.Message) response.MessageRes{
	res := response.MessageRes{Id: msg.Id, ChatId: msg.ChatId, SenderId: msg.SenderId, Content: msg.Content, TimeStamp: msg.TimeStamp}

	return res
}