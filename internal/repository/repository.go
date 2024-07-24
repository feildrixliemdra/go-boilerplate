package repository

import (
	"github.com/jmoiron/sqlx"
	"go-boilerplate/internal/appcontext"
)

type Repository struct {
	UserRepository IUserRepository
}

type Option struct {
	appcontext.Option
	DB *sqlx.DB
}

func InitiateRepository(opt Option) *Repository {
	return &Repository{
		UserRepository: NewUserRepository(opt),
	}
}
