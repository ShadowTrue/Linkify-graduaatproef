package game_repo

import (
	"backend/src/models"

	"github.com/google/uuid"
)

type GameRepo interface {
	GetGameById(gameId uuid.UUID) *models.Game
	GetAllGames() *[]models.Game
	CreateGame(game models.Game)
	UpdateUserGamesList(game []models.Game)
	UpdateGame (game models.Game)
	DeleteGame (gameId uuid.UUID)
}