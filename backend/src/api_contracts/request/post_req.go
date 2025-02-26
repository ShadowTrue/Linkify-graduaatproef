package request

import "github.com/google/uuid"

type PostReq struct {
	content string
	image   string
	sender  uuid.UUID
}