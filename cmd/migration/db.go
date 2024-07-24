package migration

import (
	"go-boilerplate/internal/bootstrap"
	"go-boilerplate/pkg/dbmigration"
)

func MigrateDatabase() {
	cfg := bootstrap.NewConfig()

	dbmigration.DatabaseMigration(cfg)
}
