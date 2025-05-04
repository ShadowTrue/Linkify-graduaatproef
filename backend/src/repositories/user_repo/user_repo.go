package user_repo

import (
	"backend/src/enums"
	"backend/src/models"

	"github.com/google/uuid"
)

type IUserRepo interface {
	GetUserByGames(games []models.Game) (*[]models.User,error)
	GetUserById(id uuid.UUID) (*models.User,error)
	GetusersByGameCategories(cat []enums.GameCategory) (*[]models.User,error)
	GetUsersByFullName(name string) (*[]models.User,error)
	CreateUser(usr *models.User) error
	UpdateUser(usr *models.User) error
	DeleteUser(id uuid.UUID) error
}