package routes

import (
	"freecreate/internal/config"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/resend/resend-go/v2"
	"github.com/valkey-io/valkey-go"
)

func MobileRouter(pgCore *pgxpool.Pool, pgContentPools config.PgContentPools, pgCoreQueries config.PgCoreQueries, pgContentQueries config.PgContentQueries, valkeyClient valkey.Client, resendClient *resend.Client) *chi.Mux {

	router := chi.NewRouter()

	return router
}
