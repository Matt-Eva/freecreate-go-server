package web_page_handlers

import (
	"html/template"
	"net/http"
)

func HomePageHandler(homeTmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type CardContent struct {
			CardTitle string
			CardDescription string
		}

		type PageData struct {
			Title    string
			LoggedIn bool
			UserIsAdult bool
			LoadedContent []CardContent
		}

		cardContent := make([]CardContent, 0, 100)

		for i := 0; i < 100; i++{
			cardContent = append(cardContent, CardContent{CardTitle: "welcome home", CardDescription: "a heartwarming tale"})
		}

		pageData := PageData{
			Title:    "home",
			LoggedIn: false,
			UserIsAdult: true,
			LoadedContent: cardContent,
		}
		homeTmpl.ExecuteTemplate(w, "home", pageData)
	}
}
