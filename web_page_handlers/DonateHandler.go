package web_page_handlers

import (
	"html/template"
	"net/http"
)

func DonateHandler(donateTmpl *template.Template) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		type PageData struct{
			LoggedIn bool
		}

		pageData := PageData {
			LoggedIn: false,
		}

		donateTmpl.ExecuteTemplate(w, "donate", pageData)
	}
}