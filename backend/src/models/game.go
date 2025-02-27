package models

import (
	"backend/src/enums"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Game struct {
	Id          uuid.UUID
	Image       string
	Name        string
	ReleaseDate time.Time
	Category    enums.GameCategory
}

func NewGame(id uuid.UUID, image, name string, releaseDate time.Time, category enums.GameCategory) (*Game, []error) {

	newGame := &Game{id, image, name, releaseDate, category}

	errs := validateGame(newGame)

	if errs != nil {
		return nil, errs
	}
	return newGame, nil
}

func validateGame(game *Game) []error {

	var errs []error

	if strings.TrimSpace(game.Image) == "" {
		errs = append(errs, errors.New("image is required"))
	}
	if strings.TrimSpace(game.Name) == "" {
		errs = append(errs, errors.New("name is required"))
	}
	if game.ReleaseDate.IsZero() {
		errs = append(errs, errors.New("releaseDate is required"))
	}

	if len(errs) > 0 {
		return errs
	}
	return nil

}
