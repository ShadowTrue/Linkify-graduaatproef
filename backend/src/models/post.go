package models

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	Id        uuid.UUID
	Sender    uuid.UUID
	Content   string
	Image     string
	TimeStamp time.Time
}

func NewPost(id, sender uuid.UUID, content, image string, timeStamp time.Time) (*Post, []error) {
	if id == uuid.Nil {
		id = uuid.New()
	}
	if timeStamp.IsZero() {
		timeStamp = time.Now().UTC()
	}
	newPost := &Post{id, sender, content, image, timeStamp}

	errs := validatePost(newPost)

	if errs != nil {
		return nil, errs
	}
	return newPost, nil
}

func validatePost(pst *Post) []error {
	var errs []error

	if pst.Sender == uuid.Nil {
		errs = append(errs, errors.New("sender is required"))
	}
	if strings.TrimSpace(pst.Content) == "" && strings.TrimSpace(pst.Image) == "" {
		errs = append(errs, errors.New("content or image is required"))
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
