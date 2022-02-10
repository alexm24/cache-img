package service

import (
	"github.com/alexm24/cache-img/internal/models"
	"github.com/alexm24/cache-img/internal/repository"
	"github.com/alexm24/cache-img/internal/service/postgres"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: postgres.NewAuth(repos.Authorization),
	}
}
