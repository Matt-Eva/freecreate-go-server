package routes

import (
	"freecreate/internal/config"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/resend/resend-go/v2"
	"github.com/valkey-io/valkey-go"
)

func CreateRouter(webSessionStore *sessions.CookieStore, pgCore *pgxpool.Pool, pgContentPools config.PgContentPools, pgCoreQueries config.PgCoreQueries, pgContentQueries config.PgContentQueries, valkeyClient valkey.Client, resendClient *resend.Client) *chi.Mux {

	router := chi.NewRouter()

	ConfigureWebRouter(router, webSessionStore, valkeyClient, pgCore, pgContentPools, pgCoreQueries, pgContentQueries, resendClient)

	router.Mount("/desktop-api", DesktopRouter(pgCore, pgContentPools, pgCoreQueries, pgContentQueries, valkeyClient, resendClient))

	router.Mount("/mobile-api", MobileRouter(pgCore, pgContentPools, pgCoreQueries, pgContentQueries, valkeyClient, resendClient))

	return router
}
