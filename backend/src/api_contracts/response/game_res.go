package response

import (
	"backend/src/enums"
	"time"

	"github.com/google/uuid"
)

type GameRes struct{
	id uuid.UUID
	image string
	name string
	releaseDate time.Time
	category enums.GameCategory
}