package web_page_handlers

import (
	"freecreate/auth"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func HomePageHandler(homeTmpl *template.Template, sessionStore *sessions.CookieStore, valkeyClient valkey.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		_, userId, _ := auth.GetUser(ctx, sessionStore, valkeyClient, w, r)
		loggedIn := false
		loggedInClass := "logged_out"
		if userId != 0 {
			loggedIn = true
			loggedInClass = "logged_in"
		}

		writingType := chi.URLParam(r, "writing_type")

		type CardContent struct{
			Title string
			Author string
			WritingType string
			Genres []string
			Description string
		}

		type CardContentCategory struct {
			Category string
			Content []CardContent
		}

		type PageData struct {
			LoggedIn bool
			LoggedInClass string
			WritingType string
			CardContentCategories []CardContentCategory
		}

		cardContentCategories := []CardContentCategory{}
		
		for i := 0; i < 50; i++{
			cardContentCategory := CardContentCategory{
				Category: "Fiction",
				Content: []CardContent{},
			}

			for i:= 0; i <50; i++{
				cardContent := CardContent{
					Title: "test",
					Author: "test",
					WritingType: "Short Story",
					Genres: []string{"Drama", "Fantasy", "Romance"},
					Description: "This is a scintillating description! We want it to be about 1000 characters in length. How long do you think it will take to reach 1000 character? And if we made it even further beyond space and time? How about if we got to 200 characters? Huh? About 250? Maybe we should consider bumping it up to 300?",
				}
				cardContentCategory.Content = append(cardContentCategory.Content, cardContent)
			}
			
			cardContentCategories = append(cardContentCategories, cardContentCategory)
		}

		pageData := PageData{
			LoggedIn: loggedIn,
			LoggedInClass: loggedInClass,
			WritingType: writingType,
			CardContentCategories: cardContentCategories,

		}


		homeTmpl.ExecuteTemplate(w, "home", pageData)
	}
}
