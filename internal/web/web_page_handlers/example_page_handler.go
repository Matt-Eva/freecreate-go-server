package web_page_handlers

import (
	"fmt"
	"freecreate/internal/config"
	"freecreate/internal/query_handlers"
	"freecreate/internal/web/web_auth"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valkey-io/valkey-go"
)

// Page handler functions sould always be named "PageNamePageHandler".
// Functions must be capitalized in order to be exported from the package, as per Go's
// package export rules.

// These parent, wrapper functions should always return an "http.HandlerFunc"
// to follow the convention established by Go's core net/http package.
// We return a function that matches the function signature of "http.HandlerFunc"
// and make use of a closure to pass in the params from the parent wrapper func.

// The parameters we're defining in this parent wrapper function are common parameters you will need in
// many page handlers.
// The exact parameters you will need will be defined by the queries necessary to load the page itself.
// You will pretty much always need to pass in sessionStore and valkeyClient in order to handle user
// authentication and session management.

func ExamplePageHandler(template *template.Template, sessionStore *sessions.CookieStore, valkeyClient valkey.Client, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries)http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		// First, we want to get the context of this particular request.
		// This will be used as the context for this request lifecycle.
		// For more on context, please refer to the code guide.
		 ctx := r.Context()

		//  Next, we want to check authentication of our user.
		userId, _ := web_auth.CheckAuthentication(ctx, sessionStore, valkeyClient, w, r)
		
		// for webpages that are only available to an authorized user, we want to do the following:
		// I have this commented out for now, but you can comment it in and comment out the below section
		// To test it out.

		// if userId == 0 {
		// 	http.Redirect(w, r, "/login", 303)
		// 	return
		// }

		// loggedIn := true
		// loggedInClass := "logged_in"

		// for webpages that are available to any user, we want to do the following:

		loggedIn := false
		loggedInClass := "logged_out"

		if userId != 0{
			loggedIn = true
			loggedInClass = "logged_in"
		}

		// call the requisite query handler to load the data for this page
		queryResult, queryErr := query_handlers.HandleExampleQuery(ctx, valkeyClient, pgCore, pgCoreQueries)
		if queryErr != nil {
			// if there was an error loading the data, render the standard error page instead.
			ErrorPageHandler(template, w, queryErr.Message, loggedIn,loggedInClass)
			return
		}

		// define the page data for this page
		// all page data structs will have a LoggedIn field and LoggedInClass field
		// These are defined in the UniversalPageData struct in the structs.go file in this package.
		type PageData struct {
			UniversalPageData
			PageValues query_handlers.ExampleQueryReturnValues
		}

		pageData := PageData {
			UniversalPageData: UniversalPageData {
				LoggedIn: loggedIn,
				LoggedInClass: loggedInClass,
				CsrfToken: csrf.TemplateField(r),
			},
			PageValues: queryResult,
		}

		fmt.Println(pageData.LoggedIn)
		fmt.Println(pageData.PageValues.ExampleParams)

		template.ExecuteTemplate(w, "example_page", pageData)
	}
}