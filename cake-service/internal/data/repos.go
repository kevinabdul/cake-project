package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Repos struct {
	Cakes CakeRepository
}

func NewRepos(db *sql.DB) Repos {
	return Repos{
		Cakes: NewDefaultCakeRepository(db),
	}
}
