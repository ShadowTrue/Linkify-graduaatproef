package request

import "github.com/google/uuid"

type ChatReq struct {
	Id uuid.UUID
	Participants []uuid.UUID
	Admins []uuid.UUID
}