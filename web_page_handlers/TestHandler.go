package web_page_handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
)

func TestHandler (testTempl *template.Template) http.HandlerFunc{

	return func (w http.ResponseWriter, r *http.Request){
		fmt.Println(r.Method)
		type PageData struct{
			LoggedIn bool
			CSRFToken template.HTML
		}

		pageData := PageData{
			LoggedIn: true,
			CSRFToken: csrf.TemplateField(r),
		}
		fmt.Println(pageData.CSRFToken)

		testTempl.ExecuteTemplate(w, "test", pageData)
	}
}