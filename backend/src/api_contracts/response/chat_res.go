package response

import "github.com/google/uuid"

type ChatRes struct {
	id uuid.UUID
	participants []uuid.UUID
	admins []uuid.UUID
	
}