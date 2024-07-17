package appcontext

import (
	"github.com/jmoiron/sqlx"
	"go-boilerplate/internal/repository"
)

func (appCtx *AppContext) InitiateRepository(dbConn *sqlx.DB) *repository.Repository {

	exampleRepository := repository.NewExampleRepository()

	return &repository.Repository{
		ExampleRepository: exampleRepository,
	}
}
