package mappers

import (
	"backend/src/api_contracts/response"
	"backend/src/models"
)

func MapFriendToRes(fr models.Friend) response.FriendRes{
	res := response.FriendRes{
		Id: fr.Id,
		ProfilePicture: fr.ProfilePicture, 
		FirstName: fr.FirstName, 
		LastName: fr.LastName, 
		FullName: fr.FullName, 
		BirthDate:  fr.BirthDate}

		return res
		
}