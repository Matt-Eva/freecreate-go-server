package web_page_handlers

import (
	"fmt"
	"freecreate/internal/lib/logger"
	"freecreate/internal/web/web_auth"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func SearchPageHandler(searchTmpl *template.Template, sessionStore *sessions.CookieStore, valkeyClient valkey.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userId, _ := web_auth.CheckAuthentication(ctx, sessionStore, valkeyClient, w, r)
		loggedIn := false
		loggedInClass := "logged_out"
		if userId != 0 {
			loggedIn = true
			loggedInClass = "logged_in"
		}

		query := r.URL.Query()
		fmt.Println(query)
		searchParams := query["search"]
		tags := query["tags"]
		fmt.Println(tags)

		type PageData struct {
			Query string
			UniversalPageData
		}

		pageData := PageData{
			UniversalPageData: UniversalPageData{
				LoggedIn:      loggedIn,
				LoggedInClass: loggedInClass,
				CsrfToken:     csrf.TemplateField(r),
			},
		}

		if len(searchParams) > 0 {
			pageData.Query = searchParams[0]
		} else {
			pageData.Query = ""
		}

		err := searchTmpl.ExecuteTemplate(w, "search_page", pageData)
		if err != nil {
			logger.Log(err)
		}
	}
}
