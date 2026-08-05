package main

import (
	"freecreate/config"
	"freecreate/routes"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/resend/resend-go/v2"
	"github.com/valkey-io/valkey-go"
)

func CreateRouter(sessionStore *sessions.CookieStore, pgxPools config.PgxPools, pgCoreQueries config.PgCoreQueries, valkeyClient valkey.Client, resendClient *resend.Client) *chi.Mux {

	router := chi.NewRouter()

	routes.ConfigureWebRouter(router, sessionStore, valkeyClient, pgxPools, pgCoreQueries, resendClient)

	router.Mount("/electron-api", routes.ElectronRouter())

	router.Mount("/react-native-api", routes.ReactNativeRouter())

	return router
}
