package web_page_handlers

import (
	web_page_utils "freecreate/web_page_handlers/utils"
	"html/template"
	"net/http"
)

func HomePageHandler(homeTmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		preloadLinks := []string{"/static/globals.css", "/static/header_component.css"}
		web_page_utils.HandlePreloadLinks(w, preloadLinks)

		type CardContent struct {
			CardTitle       string
			CardDescription string
		}

		type PageData struct {
			Title         string
			LoggedIn      bool
			UserIsAdult   bool
			LoadedContent []CardContent
		}

		cardContent := make([]CardContent, 0, 100)

		for i := 0; i < 100; i++ {
			cardContent = append(cardContent, CardContent{CardTitle: "welcome home", CardDescription: "a heartwarming tale"})
		}

		pageData := PageData{
			Title:         "home",
			LoggedIn:      false,
			UserIsAdult:   true,
			LoadedContent: cardContent,
		}
		homeTmpl.ExecuteTemplate(w, "home", pageData)
	}
}
