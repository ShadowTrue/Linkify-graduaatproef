package models

import (
	"errors"
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

func NewPost(id, sender uuid.UUID, content, image string, timeStamp time.Time) (*post, []error) {
	//TODO
	//implement this
	if id == uuid.Nil {
		id = uuid.New()
	}
	if timeStamp.IsZero() {
		timeStamp = time.Now().UTC()
	}
	newPost := &post{id, sender, content, image, timeStamp}

	errs := validatePost(newPost)

	if errs != nil {
		return nil, errs
	}
	return newPost, nil
}

func validatePost(pst *post) []error {
	var errs []error

	if pst.sender == uuid.Nil {
		errs = append(errs, errors.New("sender is required"))
	}
	if pst.content == "" && pst.image == "" {
		errs = append(errs, errors.New("content or image is required"))
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
