package web_page_handlers

import (
	"fmt"
	"freecreate/config"
	pg_core_queries "freecreate/db/pg_core/queries"
	"freecreate/logger"
	web_page_utils "freecreate/web_page_handlers/utils"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/jackc/pgx/v5/pgxpool"
)

func LoginPageHandler(loginTmpl *template.Template, pgxCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		switch r.Method {
		case "GET":
			handleLoginPageGet(loginTmpl, w, r)
		case "POST":
			handleLoginPagePost(loginTmpl, w, r, pgxCore, pgCoreQueries)
		default:
			web_page_utils.HandleInvalidWebpageRequestMethod(w, r.Method)
		}
		
	}
}

func renderLoginPage(loginTmpl *template.Template, w http.ResponseWriter, r *http.Request, errors []string){
	
		type PageData struct {
			LoggedIn bool
			RequestMethod string
			CSRFToken template.HTML
			Errors []string
		}

		pageData := PageData{
			LoggedIn: false,
			RequestMethod: r.Method,
			CSRFToken: csrf.TemplateField(r),
			Errors: errors,
		}

		loginTmpl.ExecuteTemplate(w, "login_page", pageData)
}

func handleLoginPageGet(loginTmpl *template.Template, w http.ResponseWriter, r *http.Request){
	renderLoginPage(loginTmpl, w, r, []string{})
}

func handleLoginPagePost(loginTmpl *template.Template, w http.ResponseWriter, r *http.Request, pgxCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries){
	var errs []string;

	formAction, formActionErr := web_page_utils.GetFormAction(r);
	if formActionErr != nil {
		logger.Log(formActionErr)
		errs = append(errs, formActionErr.Error())
		renderLoginPage(loginTmpl, w, r, errs)
		return
	}
	fmt.Println(formAction)

	userId, getUserErr := pg_core_queries.GetUserByEmail(pgCoreQueries, pgxCore, "")
	if getUserErr != nil {
		logger.Log(getUserErr)
		errs = append(errs, getUserErr.Error())
		renderLoginPage(loginTmpl, w, r, errs)
		return
	}

	fmt.Println(userId)

	renderLoginPage(loginTmpl, w, r, errs)
}
