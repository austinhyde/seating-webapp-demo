package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/austinhyde/seating/db"

	"github.com/alexflint/go-arg"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zerologadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
)

type Args struct {
	DatabaseURL string `arg:"env:DB_URL,required"`

	HostPort string `arg:"positional,required"`
}

func main() {
	args := Args{}
	arg.MustParse(&args)

	logger := zerolog.New(zerolog.NewConsoleWriter())
	log := &logger

	mainCtx := context.Background()

	dbcfg, err := pgxpool.ParseConfig(args.DatabaseURL)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to parse db config")
	}
	dbcfg.ConnConfig.Logger = zerologadapter.NewLogger(logger)
	dbcfg.ConnConfig.LogLevel = pgx.LogLevelTrace

	conn, err := pgxpool.Connect(mainCtx, args.DatabaseURL)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to connect to database")
	}
	defer conn.Close()

	err = db.ApplyMigrations(mainCtx, log, conn)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to apply migrations")
	}

	http.HandleFunc("/", HelloServer)
	_ = http.ListenAndServe(args.HostPort, nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world, %s!", r.URL.Path[1:])
}
