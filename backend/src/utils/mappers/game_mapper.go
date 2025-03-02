package mappers

import (
	"backend/src/api_contracts/response"
	"backend/src/models"
)

func MapGameToRes(game *models.Game)  *response.GameRes{
	res := response.GameRes{Id: game.Id, 
		Image: game.Image, 
		Name: game.Name, 
		ReleaseDate: game.ReleaseDate , 
		Category: game.Category}

		return &res
}
