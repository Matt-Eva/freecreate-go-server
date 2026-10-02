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

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valkey-io/valkey-go"
)

func EditWritingPageHandler(templates template.Template, sessionStore *sessions.CookieStore, valkeyClient valkey.Client, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// startTime := time.Now()

		ctx := r.Context()

		userId, _ := web_auth.CheckAuthentication(ctx, sessionStore, valkeyClient, w, r)
		if userId == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		writingUuidParam := chi.URLParam(r,"writing_uuid")

		writingUuid, uuidParseErr := uuid.Parse(writingUuidParam)
		if uuidParseErr != nil {
			logger.Log(uuidParseErr)
			http.Error(w, api_error.InteralServerErrorMessage, http.StatusInternalServerError)
			return
		}

		getMyWritingParams := pg_core_types.GetMyWritingParams{
			UUID: writingUuid,
			UserId: userId,
		}

		myWriting, getMyWritingErr := query_handlers.HandleGetMyWriting(ctx, pgCore, pgCoreQueries, getMyWritingParams)
		if getMyWritingErr != nil {
			http.Error(w, getMyWritingErr.Message, getMyWritingErr.Code)
			return
		}

		getMyCreatorsParams := pg_core_types.GetMyCreatorsParams{
			UserId: userId,
		}

		myCreators, getMyCreatorsErr := query_handlers.HandleGetMyCreators(ctx, pgCore, pgCoreQueries, getMyCreatorsParams)
		if getMyCreatorsErr != nil {
			http.Error(w, getMyCreatorsErr.Message, getMyCreatorsErr.Code)
			return
		}

		var currentCreator pg_core_types.MyCreatorsStruct

		for i := 0; i < len(myCreators); i++{
			creator := myCreators[i]
			if creator.ID == myWriting.CreatorId{
				currentCreator = creator
				break
			}
		}

		type PageData struct {
			UniversalPageData
			MyWriting pg_core_types.MyWriting
			MyCreators []pg_core_types.MyCreatorsStruct
			CurrentCreator pg_core_types.MyCreatorsStruct
		}

		pageData := PageData{
			UniversalPageData: UniversalPageData{
				LoggedIn:      true,
				LoggedInClass: "logged_in",
				CsrfToken:     csrf.TemplateField(r),
			},
			MyWriting: myWriting,
			MyCreators: myCreators,
			CurrentCreator: currentCreator,
		}

		fmt.Println(pageData)

		err := templates.ExecuteTemplate(w, "edit_writing_page", pageData)
		if err != nil {
			logger.Log(err)
		}

		// duration := time.Since(startTime)

		// fmt.Println(duration)
	}
}
