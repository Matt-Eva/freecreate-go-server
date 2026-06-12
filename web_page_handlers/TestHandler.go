package web_page_handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
)

func TestHandler (testTempl *template.Template) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		requestMethod := r.Method
		switch requestMethod {
		case "GET":
			handleGet(testTempl, w, r)
		case "POST":
			handlePost(testTempl, w, r)
		}
	}
}

func handleGet(testTmpl *template.Template, w http.ResponseWriter, r *http.Request){
	type PageData struct{
			LoggedIn bool
			CSRFToken template.HTML
		}
	pageData := PageData{
		LoggedIn: true,
		CSRFToken: csrf.TemplateField(r),
	}
	testTmpl.ExecuteTemplate(w, "test", pageData)

}

func handlePost(testTmpl *template.Template, w http.ResponseWriter, r *http.Request){
	contentType := r.Header.Get("Content-Type")
		switch contentType {
		case "application/json":
			handleJSONPost(w, r)
		default:
			handleFormPost(testTmpl, w, r)
		}
}

func handleFormPost(testTmpl *template.Template, w http.ResponseWriter, r *http.Request){
	fmt.Println("handling form submission")
	type PageData struct{
			LoggedIn bool
			CSRFToken template.HTML
		}
	pageData := PageData{
		LoggedIn: true,
		CSRFToken: csrf.TemplateField(r),
	}
	testTmpl.ExecuteTemplate(w, "test", pageData)
}

func handleJSONPost(w http.ResponseWriter, r *http.Request){
	fmt.Println("handling json")
	w.Header().Set("Content-Type", "application/json")
	type JSONResponse struct {
		Message string `json:"message"`
	}
	res := JSONResponse {
		Message: "successful json post!",
	}
	json.NewEncoder(w).Encode(res)
}	