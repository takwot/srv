package service

import (
	todo "github.com/takwot/srv"
	"github.com/takwot/srv/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type Film interface {
}

type Films interface {
}

type Service struct {
	Authorization
	Film
	Films
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
	}
}
