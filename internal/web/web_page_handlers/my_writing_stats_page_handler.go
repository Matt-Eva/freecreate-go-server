package web_page_handlers

import (
	"html/template"
	"net/http"
)

func MyWritingStatsPageHandler(templates *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
