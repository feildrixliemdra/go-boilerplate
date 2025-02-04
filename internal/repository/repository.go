package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	UserRepository IUserRepository
}

type Option struct {
	DB *sqlx.DB
}

func InitiateRepository(opt Option) *Repository {
	return &Repository{
		UserRepository: NewUserRepository(opt),
	}
}
