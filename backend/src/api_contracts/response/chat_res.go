package response

import "github.com/google/uuid"

type ChatRes struct {
	Id uuid.UUID
	Participants []uuid.UUID
	Admins []uuid.UUID
	
}