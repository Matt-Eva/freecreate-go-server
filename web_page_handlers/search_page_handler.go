package web_page_handlers

import (
	"html/template"
	"net/http"
)

func SearchPageHandler(searchTmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		searchParams := query["search"]

		type PageData struct {
			Query string
		}

		var pageData PageData

		if len(searchParams) > 0 {
			pageData = PageData{
				Query: searchParams[0],
			}
		} else {
			pageData = PageData{
				Query: "",
			}
		}

		searchTmpl.ExecuteTemplate(w, "search_page", pageData)
	}
}
