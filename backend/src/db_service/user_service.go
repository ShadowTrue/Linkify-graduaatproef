package dbservice

import (
	"backend/src/enums"
	"backend/src/models"
	"backend/src/utils/context"
	"errors"
	"log"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
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
func (us *UserService) GetUserByGames(games []models.Game,page int,pageSize int) (*[]models.User , error){
	ctx,cancel := context.TimeoutCtx()
	defer cancel()

	//TODO After implementing the game service

	
}

func (us *UserService) GetUserById(id uuid.UUID) (*models.User,error){
	ctx,cancel := context.TimeoutCtx()
	defer cancel()
	
	filter := bson.D{{Key: "_id", Value: id}}

	usr := &models.User{}

	err := us.collection.FindOne(ctx,filter).Decode(usr)

	if(err == mongo.ErrNoDocuments){
		err = errors.New("No user found")
	}

	if(err != nil){
		log.Println(err)
	}

	return usr,err

	}
func (us *UserService) GetusersByGameCategories(cat []enums.GameCategory) (*[]models.User,error){
	//TODO after implementing the game service
	}
func (us *UserService) GetUsersByFullName(name string) (*[]models.User,error){

	}
func (us *UserService) CreateUser(usr *models.User) error {

	 }
func (us *UserService) UpdateUser(usr *models.User) error {

	}
func (us *UserService) DeleteUser(id uuid.UUID) error {
		
	}


