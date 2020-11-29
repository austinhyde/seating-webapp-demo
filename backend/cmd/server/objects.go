package main

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zerologadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
)

func getLogger() *zerolog.Logger {
	logger := zerolog.New(zerolog.NewConsoleWriter()).With().Timestamp().Logger()
	return &logger
}

func getDatabaseConnection(ctx context.Context, log *zerolog.Logger, url string) *pgxpool.Pool {
	dbcfg, err := pgxpool.ParseConfig(url)
	maybeFatal(err, "Unable to parse db config")

	dbcfg.ConnConfig.Logger = zerologadapter.NewLogger(*log)
	dbcfg.ConnConfig.LogLevel = pgx.LogLevelTrace

	conn, err := pgxpool.ConnectConfig(ctx, dbcfg)
	maybeFatal(err, "Unable to connect to database")

	return conn
}
