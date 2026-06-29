package web_page_handlers

import (
	web_page_utils "freecreate/web_page_handlers/utils"
	"html/template"
	"net/http"
)

func AboutPageHandler(aboutTmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		preloadLinks := []string{"/static/globals.css", "/static/header_component.css"}
		web_page_utils.HandlePreloadLinks(w, preloadLinks)

		// time.Sleep(300 * time.Millisecond)

		type PageData struct {
			Title    string
			LoggedIn bool
		}

		pageData := PageData{
			Title:    "about",
			LoggedIn: false,
		}

		aboutTmpl.ExecuteTemplate(w, "about_page", pageData)
	}
}

func write103Header(w http.ResponseWriter) {
	w.Header().Add("Link", "</static/globals.css>; rel=preload; as=style")
	w.Header().Add("Link", "</static/header_component.css>; rel=preload; as=style")
	w.WriteHeader(103)
}
