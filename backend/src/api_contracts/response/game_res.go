package response

import (
	"backend/src/enums"
	"time"

	"github.com/google/uuid"
)

type GameRes struct{
	Id uuid.UUID
	Image string
	Name string
	ReleaseDate time.Time
	Category enums.GameCategory
}