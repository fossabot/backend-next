package infra

import (
	"context"
	"database/sql"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bun/extra/bunotel"

	"exusiai.dev/backend-next/internal/app/appconfig"
)

func Postgres(conf *appconfig.Config) (*bun.DB, error) {
	// Open a Postgres database.
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(conf.PostgresDSN), pgdriver.WithApplicationName("penguin-backend")))

	// Create a Bun db on top of it.
	db := bun.NewDB(pgdb, pgdialect.New())
	if conf.DevMode {
		db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithEnabled(true), bundebug.WithVerbose(conf.BunDebugVerbose)))
	}
	if conf.TracingEnabled {
		db.AddQueryHook(bunotel.NewQueryHook(bunotel.WithDBName("penguin-postgres")))
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Error().Err(err).Msg("infra: postgres: failed to ping database")
		return nil, err
	}

	pgdb.SetMaxOpenConns(conf.PostgresMaxOpenConns)
	pgdb.SetMaxIdleConns(conf.PostgresMaxIdleConns)
	pgdb.SetConnMaxLifetime(conf.PostgresConnMaxLifeTime)
	pgdb.SetConnMaxIdleTime(conf.PostgresConnMaxIdleTime)

	return db, nil
}
