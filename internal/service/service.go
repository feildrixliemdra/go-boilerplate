package service

import (
	"go-boilerplate/internal/config"
	"go-boilerplate/internal/repository"
)

type Service struct {
	UserService IUserService
}

type Option struct {
	Repository *repository.Repository
}

func InitiateService(cfg *config.Config, repository *repository.Repository) *Service {
	return &Service{
		UserService: NewUserService(repository.UserRepository),
	}
}
