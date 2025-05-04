package friend_repo

import (
	"backend/src/models"

	"github.com/google/uuid"
)

type FriendRepo interface {
	GetFriendById(friendId models.Friend) (*models.Friend, error)
	GetAllFriendByUserId(userId uuid.UUID) (*[]models.Friend,error)
	AddFriend(friend models.Friend) error
	UpdateFriendStatus(friend models.Friend) error
	DeleteFriend(friendId uuid.UUID) error
}