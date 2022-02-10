package repository

import (
	"github.com/alexm24/cache-img/internal/models"
	"github.com/alexm24/cache-img/internal/repository/postgres"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: postgres.NewAuth(db),
	}
}
