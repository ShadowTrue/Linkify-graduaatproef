package mappers

import (
	"backend/src/api_contracts/request"
	"backend/src/api_contracts/response"
	"backend/src/models"
	"time"

	"github.com/google/uuid"
)

func MapPostToStruct(req request.PostReq) (*models.Post,[]error){
	post,err := models.NewPost(uuid.Nil,req.Sender,req.Content,req.Image,time.Time{})

	return post,err
}

func MapPostToRes(post models.Post) *response.PostRes{
	res := response.PostRes{Id: post.Id, Sender: post.Sender , Content: post.Content, Image: post.Image, TimeStamp: post.TimeStamp}

	return &res
}