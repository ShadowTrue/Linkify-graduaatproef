package mappers

import (
	"backend/src/api_contracts/request"
	"backend/src/models"
)

func MapUserToStruct(req request.UserReq) (*models.User,[]error){
	 usr,errors := models.NewUser(req.Id,req.ProfilePicture,req.Email,req.Username,req.FirstName,req.LastName,req.Birthday,req.Country,req.CreatedOn)
	return usr,errors
}

func MapUserToRes(usr models.User) (
	//TODO
	//Remeber to make all params in contracts public
)

