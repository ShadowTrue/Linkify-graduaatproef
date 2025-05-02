package friend_repo

import (
	"backend/src/models"

	"github.com/google/uuid"
)

type FriendRepo interface {
	GetFriendById(friendId models.Friend) *models.Friend
	GetAllFriendByUserId(userId uuid.UUID) *[]models.Friend
	AddFriend(friend models.Friend)
	UpdateFriendStatus(friend models.Friend)
	DeleteFriend(friendId uuid.UUID)
}