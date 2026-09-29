package web_page_handlers

import (
	"fmt"
	"freecreate/internal/config"
	pg_core_types "freecreate/internal/db/pg_core/types"
	"freecreate/internal/lib/api_error"
	"freecreate/internal/lib/logger"
	"freecreate/internal/query_handlers"
	"freecreate/internal/web/web_auth"
	"html/template"
	"net/http"

	"github.com/google/uuid"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valkey-io/valkey-go"
)

func MyWritingPageHandler(template *template.Template, sessionStore *sessions.CookieStore, valkeyClient valkey.Client, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		userId, _ := web_auth.CheckAuthentication(ctx, sessionStore, valkeyClient, w, r)
		if userId == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		writingUUIDParam := chi.URLParam(r, "writing_uuid")

		writingUUID, uuidErr := uuid.Parse(writingUUIDParam)
		if uuidErr != nil {
			logger.Log(uuidErr)
			http.Error(w, api_error.InteralServerErrorMessage, http.StatusInternalServerError)
			return
		}

		getWritingParams := pg_core_types.GetMyWritingParams{
			UserId: userId,
			UUID:   writingUUID,
		}

		myWriting, queryErr := query_handlers.HandleGetMyWriting(ctx, pgCore, pgCoreQueries, getWritingParams)
		if queryErr != nil {
			http.Error(w, queryErr.Message, queryErr.Code)
			return
		}

		type PageData struct {
			UniversalPageData
			MyWriting pg_core_types.MyWriting
		}

		pageData := PageData{
			UniversalPageData: UniversalPageData{
				LoggedIn:      true,
				LoggedInClass: "logged_in",
				CsrfToken:     csrf.TemplateField(r),
			},
			MyWriting: myWriting,
		}

		fmt.Println(pageData.MyWriting)

		err := template.ExecuteTemplate(w, "my_writing_page", pageData)
		if err != nil {
			logger.Log(err)
		}
	}
}
