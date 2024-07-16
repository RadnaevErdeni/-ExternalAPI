package service

import (
	"tt/repository"
	"tt/testtask"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetByPas(serie, number int) (testtask.Users, error) {
	return s.repo.GetByPas(serie, number)
}
