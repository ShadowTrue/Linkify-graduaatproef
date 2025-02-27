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
	res := response.UserRes{
		Id: usr.Id,
		ProfilePicture: usr.ProfilePicture,
		Email: usr.Email,
		Username: usr.Username,
		FirstName: usr.FirstName,
		LastName: usr.LastName,
		Birthday: usr.Birthday,
		Country: usr.Country,
		CreatedOn: usr.CreatedOn}
	return res
}

	


