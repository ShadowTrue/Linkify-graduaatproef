package post_repo

import (
	"backend/src/models"

	"github.com/google/uuid"
)

type PostRepo interface {
	GetPostById(postId uuid.UUID)
	GetUserPosts(userId uuid.UUID) []models.Post
	GetFriendsPosts(userId uuid.UUID) []models.Post
	GetLatestPostsByRandomUsers() []models.Post
	GetPostsByUser(userId uuid.UUID) []models.Post
	CreatePost(post models.Post)
	UpdatePost(post models.Post)
	DeletePost(postId models.Post)
}