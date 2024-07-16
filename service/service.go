package service

import (
	"tt/repository"
	"tt/testtask"
)

type User interface {
	GetByPas(serie, number int) (testtask.Users, error)
}

type Service struct {
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
	}
}
