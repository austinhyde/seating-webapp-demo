package main

import (
	"context"
	"net/http"

	"github.com/alexflint/go-arg"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"

	"github.com/austinhyde/seating/db"
)

const assetsDir = "static"
const indexHTML = assetsDir + "/index.html"

var args struct {
	DatabaseURL string `arg:"env:DB_URL,required"`
	HostPort    string `arg:"positional,required"`
}
var log *zerolog.Logger
var conn *pgxpool.Pool
var mainCtx context.Context

func main() {
	arg.MustParse(&args)

	mainCtx = context.Background()
	log = getLogger()

	conn = getDatabaseConnection(mainCtx, log, args.DatabaseURL)
	defer conn.Close()

	err := db.ApplyMigrations(mainCtx, log, conn)
	maybeFatal(err, "Unable to apply migrations")

	startServer()
}

func maybeFatal(err error, msg string) {
	if err != nil {
		log.Fatal().Err(err).Msg(msg)
	}
}

func startServer() {
	m := mux.NewRouter()
	m.Use(httpLogger(log))
	// m.PathPrefix("/api").Handler(
	// 	http.StripPrefix(
	// 		"/api",
	// 		getAPIRoutes(),
	// 	),
	// )
	m.PathPrefix("/static").Handler(
		http.StripPrefix(
			"/static",
			http.FileServer(http.Dir(assetsDir)),
		),
	)
	m.PathPrefix("/").Handler(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, indexHTML)
		}),
	)

	s := &http.Server{
		Addr:    args.HostPort,
		Handler: m,
	}
	log.Info().Str("listen", args.HostPort).Msg("starting http server")
	log.Warn().Err(s.ListenAndServe()).Msg("http server shutdown")
}
