package web_page_handlers

import (
	"freecreate/internal/config"
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

func EditChapterPageHandler(template *template.Template, sessionStore *sessions.CookieStore, valkeyClient valkey.Client, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, pgContentPools config.PgContentPools, pgContentQueries config.PgContentQueries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userId, _ := web_auth.CheckAuthentication(ctx, sessionStore, valkeyClient, w, r)
		if userId == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		writingUuidParam := chi.URLParam(r, "writing_uuid")

		writingUuid, uuidParseErr := uuid.Parse(writingUuidParam)
		if uuidParseErr != nil {
			logger.Log(uuidParseErr)
			http.Error(w, api_error.InteralServerErrorMessage, http.StatusInternalServerError)
			return
		}

		chapterUuidParam := chi.URLParam(r, "chapter_uuid")

		chapterUuid, chapterUuidParseErr := uuid.Parse(chapterUuidParam)
		if chapterUuidParseErr != nil {
			logger.Log(chapterUuidParseErr)
			http.Error(w, api_error.InteralServerErrorMessage, http.StatusInternalServerError)
			return
		}

		_, _, _, getEditChapterErr := query_handlers.HandleGetEditChapter(ctx, pgCore, pgCoreQueries, pgContentPools, pgContentQueries, userId, writingUuid, chapterUuid)
		if getEditChapterErr != nil {
			http.Error(w, getEditChapterErr.Message, getEditChapterErr.Code)
			return
		}

		type PageData struct {
			UniversalPageData
		}

		pageData := PageData{
			UniversalPageData: UniversalPageData{
				LoggedIn:      true,
				LoggedInClass: "logged_in",
				CsrfToken:     csrf.TemplateField(r),
			},
		}

		err := template.ExecuteTemplate(w, "edit_chapter_page", pageData)
		if err != nil {
			logger.Log(err)
		}
	}
}
