package models

import (
	"backend/src/enums"
	"errors"
	"github.com/google/uuid"
	"strings"
	"time"
)

type game struct {
	id          uuid.UUID
	image       string
	name        string
	releaseDate time.Time
	category    enums.GameCategory
}

func NewGame(id uuid.UUID, image, name string, releaseDate time.Time, category enums.GameCategory) (*game, []error) {

	newGame := &game{id, image, name, releaseDate, category}

	errs := validateGame(newGame)

	if errs != nil {
		return nil, errs
	}
	return newGame, nil
}

func validateGame(game *game) []error {

	var errs []error

	if strings.TrimSpace(game.image) == "" {
		errs = append(errs, errors.New("image is required"))
	}
	if strings.TrimSpace(game.name) == "" {
		errs = append(errs, errors.New("name is required"))
	}
	if game.releaseDate.IsZero() {
		errs = append(errs, errors.New("releaseDate is required"))
	}

	if len(errs) > 0 {
		return errs
	}
	return nil

}
