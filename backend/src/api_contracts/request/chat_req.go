package request

import "github.com/google/uuid"

type ChatReq struct {
	id uuid.UUID
	participants []uuid.UUID
	admins []uuid.UUID
}