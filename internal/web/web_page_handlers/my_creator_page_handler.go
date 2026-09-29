package web_page_handlers

import (
	"freecreate/internal/config"
	pg_core_types "freecreate/internal/db/pg_core/types"
	"freecreate/internal/lib/api_error"
	"freecreate/internal/lib/logger"
	"freecreate/internal/query_handlers"
	"freecreate/internal/web/web_auth"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valkey-io/valkey-go"
)

func MyCreatorPageHandler(templates *template.Template, sessionStore *sessions.CookieStore, valkeyClient valkey.Client, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userId, _ := web_auth.CheckAuthentication(ctx, sessionStore, valkeyClient, w, r)
		if userId == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		creatorUuidParam := chi.URLParam(r, "creator_uuid")

		creatorUUID, uuidErr := uuid.Parse(creatorUuidParam)
		if uuidErr != nil {
			logger.Log(uuidErr)
			http.Error(w, api_error.InteralServerErrorMessage, http.StatusInternalServerError)
			return
		}

		getMyCreatorParams := pg_core_types.GetMyCreatorParams{
			UserId:      userId,
			CreatorUUID: creatorUUID,
		}

		myCreator, getMyCreatorErr := query_handlers.HandleGetMyCreator(ctx, pgCore, pgCoreQueries, getMyCreatorParams)
		if getMyCreatorErr != nil {
			http.Error(w, getMyCreatorErr.Message, getMyCreatorErr.Code)
			return
		}

		getMyCreatorWritingsParams := pg_core_types.GetMyCreatorWritingsParams{
			UserId:    userId,
			CreatorId: myCreator.ID,
		}

		myCreatorWritings, myCreatorWritingsErr := query_handlers.HandleGetMyCreatorWritings(ctx, pgCore, pgCoreQueries, getMyCreatorWritingsParams)
		if myCreatorWritingsErr != nil {
			http.Error(w, myCreatorWritingsErr.Message, myCreatorWritingsErr.Code)
			return
		}

		type PageData struct {
			LoggedIn          bool
			LoggedInClass     string
			MyCreator         pg_core_types.MyCreator
			MyCreatorWritings []pg_core_types.MyCreatorWritings
		}

		pageData := PageData{
			LoggedIn:          true,
			LoggedInClass:     "logged_in",
			MyCreator:         myCreator,
			MyCreatorWritings: myCreatorWritings,
		}

		err := templates.ExecuteTemplate(w, "my_creator_page", pageData)
		if err != nil {
			logger.Log(err)
		}
	}
}
