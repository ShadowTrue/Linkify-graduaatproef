package post_repo

import (
	"backend/src/models"

	"github.com/google/uuid"
)

type IPostRepo interface {
	GetPostById(postId uuid.UUID) error
	GetUserPosts(userId uuid.UUID) (*[]models.Post,error)
	GetFriendsPosts(userId uuid.UUID) (*[]models.Post,error)
	GetLatestPostsByRandomUsers() (*[]models.Post,error)
	GetPostsByUser(userId uuid.UUID) (*[]models.Post,error)
	CreatePost(post models.Post) error
	UpdatePost(post models.Post) error
	DeletePost(postId models.Post) error
}