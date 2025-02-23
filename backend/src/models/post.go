package models

import (
	"github.com/google/uuid"
	"time"
)

type post struct {
	id        uuid.UUID
	sender    uuid.UUID
	content   string
	image     string
	timeStamp time.Time
}

func NewPost(id, sender uuid.UUID) *post {
	//TODO
	//implement this
}
