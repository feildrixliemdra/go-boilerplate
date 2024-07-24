package bootstrap

import (
	"github.com/jmoiron/sqlx"
	"go-boilerplate/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitiatePostgreSQL(cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", cfg.Postgre.URL) // you can change it to mysql or any other supported sql db
	if err != nil {
		return db, err
	}

	db.SetMaxIdleConns(cfg.Postgre.MaxIdleConn)
	db.SetMaxOpenConns(cfg.Postgre.MaxOpenConn)

	return db, nil
}
