package web_page_handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"freecreate/logger"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
)

func TestPageHandler(templates *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestMethod := r.Method
		switch requestMethod {
		case "GET":
			handleTestPageGet(templates, w, r)
		case "POST":
			handleTestPagePost(templates, w, r)
		default:
			handleInvalidWebpageRequestMethod(w, requestMethod)
			return
		}
	}
}

func renderTestPage(w http.ResponseWriter, r *http.Request, templates *template.Template) {
	type testPageData struct {
		LoggedIn  bool
		CSRFToken template.HTML
	}

	pageData := testPageData{
		LoggedIn:  true,
		CSRFToken: csrf.TemplateField(r),
	}

	templates.ExecuteTemplate(w, "test", pageData)
}

func handleTestPageGet(templates *template.Template, w http.ResponseWriter, r *http.Request) {
	renderTestPage(w, r, templates)
}

func handleTestPagePost(templates *template.Template, w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	switch contentType {
	case "application/json":
		handleJSONPost(w, r)
	default:
		handleFormPost(templates, w, r)
	}
}

func handleFormPost(templates *template.Template, w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling form submission")
	parseFormErr := r.ParseForm()

	if parseFormErr != nil {
		fmt.Println(parseFormErr)
	}

	formAction := r.FormValue("form_action")
	fmt.Println(formAction)

	renderTestPage(w, r, templates)
}

func handleJSONPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling json")
	type Body struct {
		FormAction string `json:"formAction"`
	}

	var body Body

	bErr := json.NewDecoder(r.Body).Decode(&body)
	if bErr != nil {
		logger.Log(bErr)
		http.Error(w, bErr.Error(), http.StatusUnprocessableEntity)
		return
	}

	switch body.FormAction {
	case "update_profile":
		fmt.Println("updating profile")
	case "update_settings":
		fmt.Println("updating settings")
	default:
		formActionErr := errors.New("this form action is not valid for this route")
		logger.Log(formActionErr)
		http.Error(w, formActionErr.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	type JSONResponse struct {
		Message string `json:"message"`
	}
	res := JSONResponse{
		Message: "successful json post!",
	}
	json.NewEncoder(w).Encode(res)
}
