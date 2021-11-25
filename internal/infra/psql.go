package infra

import (
	"database/sql"
	"penguin-stats-v4/internal/config"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func ProvidePostgres(config *config.Config) (*bun.DB, error) {
	// Open a PostgreSQL database.
	dsn := "postgres://root:root@localhost:5432/penguin_structured?sslmode=disable"
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// Create a Bun db on top of it.
	db := bun.NewDB(pgdb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook())

	return db, nil
}
