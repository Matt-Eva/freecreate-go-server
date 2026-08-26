package routes

import (
	"freecreate/internal/config"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/resend/resend-go/v2"
	"github.com/valkey-io/valkey-go"
)

func CreateRouter(sessionStore *sessions.CookieStore, pgxPools config.PgxPools, pgCoreQueries config.PgCoreQueries, valkeyClient valkey.Client, resendClient *resend.Client) *chi.Mux {

	router := chi.NewRouter()

	ConfigureWebRouter(router, sessionStore, valkeyClient, pgxPools, pgCoreQueries, resendClient)

	router.Mount("/electron-api", ElectronRouter())

	router.Mount("/react-native-api", ReactNativeRouter())

	return router
}
