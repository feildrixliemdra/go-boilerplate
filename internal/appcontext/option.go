package appcontext

import (
	"go-boilerplate/internal/config"
	"go-boilerplate/internal/repository"
)

type Option struct {
	Config     *config.Config
	Repository *repository.Repository
}
