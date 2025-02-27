package mappers

import (
	"backend/src/api_contracts/request"
	"backend/src/api_contracts/response"
	"backend/src/models"
)

func MapUserToStruct(req request.UserReq) (*models.User,[]error){
	 usr,errs := models.NewUser(req.Id,req.ProfilePicture,req.Email,req.Username,req.FirstName,req.LastName,req.Birthday,req.Country,req.CreatedOn)
	return usr,errs
}

func MapUserToRes(usr models.User) response.UserRes{
	res := response.UserRes{usr.Id,usr.ProfilePicture,usr.Email,usr.Username,usr.FirstName,usr.LastName,usr.Birthday,usr.Country,usr.CreatedOn}
	return res
}

	


