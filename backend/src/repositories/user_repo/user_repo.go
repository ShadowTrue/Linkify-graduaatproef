package user_repo

import (
	"backend/src/enums"
	"backend/src/models"

	"github.com/google/uuid"
)

type UserRepo interface {
	GetUserByGames(games []models.Game) *[]models.User
	GetUserById(id uuid.UUID) *models.User
	GetusersByGameCategories(cat []enums.GameCategory) *[]models.User
	GetUsersByFullName(name string) *[]models.User
	CreateUser(usr *models.User)
	UpdateUser(usr *models.User)
	DeleteUser(id uuid.UUID)
}