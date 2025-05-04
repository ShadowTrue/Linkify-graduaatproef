package game_repo

import (
	"backend/src/models"

	"github.com/google/uuid"
)

type GameRepo interface {
	GetGameById(gameId uuid.UUID) (*models.Game,error)
	GetAllGames() (*[]models.Game, error)
	CreateGame(game models.Game) error
	UpdateUserGamesList(game []models.Game) error
	UpdateGame (game models.Game) error
	DeleteGame (gameId uuid.UUID) error
}