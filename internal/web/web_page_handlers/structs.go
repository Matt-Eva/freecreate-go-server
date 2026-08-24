package web_page_handlers

import "html/template"

type UniversalPageData struct{
	CsrfToken template.HTML
	LoggedIn bool
	LoggedInClass string
}