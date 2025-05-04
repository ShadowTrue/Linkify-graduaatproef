package dbservice

import (
	"backend/src/enums"
	"backend/src/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserService struct  {
	collection mongo.Collection
}


func CreateUserService(db mongo.Database) *UserService{
	return &UserService{
		collection: *db.Collection("Users"),
	}
}

//TODO Implement all methods below
func GetUserByGames(games []models.Game,page int,pageSize int) (*[]models.User , error){

}

func GetUserById(id uuid.UUID) (*models.User,error){

	}
func GetusersByGameCategories(cat []enums.GameCategory) (*[]models.User,error){

	}
func GetUsersByFullName(name string) (*[]models.User,error){

	}
func CreateUser(usr *models.User) error {

	 }
func UpdateUser(usr *models.User) error {

	}
func DeleteUser(id uuid.UUID) error {
		
	}


