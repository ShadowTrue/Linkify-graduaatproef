package request

import "github.com/google/uuid"

type PostReq struct {
	Content string
	Image   string
	Sender  uuid.UUID
}