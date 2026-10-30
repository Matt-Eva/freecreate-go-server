package web_page_handlers

import (
	"fmt"
	"freecreate/internal/lib"
	"freecreate/internal/lib/logger"
	"freecreate/internal/web/web_auth"
	"html/template"
	"net/http"
	"slices"

	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func SearchPageHandler(searchTmpl *template.Template, sessionStore *sessions.CookieStore, valkeyClient valkey.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userId, _ := web_auth.CheckAuthentication(ctx, sessionStore, valkeyClient, w, r)
		loggedIn := false
		loggedInClass := "logged_out"
		if userId != 0 {
			loggedIn = true
			loggedInClass = "logged_in"
		}

		query := r.URL.Query()

		searchType := query["searchType"]
		fmt.Println(searchType)

		writingType := query["writingType"]
		fmt.Println(writingType)

		datePosted := query["datePosted"]
		fmt.Println(datePosted)

		topics := query["topic"]

		fmt.Println(topics)

		tags := query["tag"]
		fmt.Println(tags)

		type PageData struct {
			UniversalPageData
			SearchType                   string
			WritingType                  string
			DatePosted                   string
			Topics                       []string
			Tags                         []string
			WritingSearchTypes           []string
			AllWritingTopics             []string
			FictionTopics                []string
			EssayAndBlogTopics           []string
			MemoirAndAutobiographyTopics []string
			PoetryTopics                 []string
		}

		var allWritingTopics []string

		for _, val := range lib.WritingTopics.Fiction {
			allWritingTopics = append(allWritingTopics, val)
		}

		for _, val := range lib.WritingTopics.EssaysAndBlogs {
			allWritingTopics = append(allWritingTopics, val)
		}

		for _, val := range lib.WritingTopics.MemoirAndAutobiography {
			allWritingTopics = append(allWritingTopics, val)
		}

		for _, val := range lib.WritingTopics.Poetry {
			allWritingTopics = append(allWritingTopics, val)
		}

		slices.Sort(allWritingTopics)
		allWritingTopics = slices.Compact(allWritingTopics)
		allWritingTopics = slices.DeleteFunc(allWritingTopics, func(topic string)bool{
			return topic == "No Genre" || topic == "No Topic"
		})

		pageData := PageData{
			UniversalPageData: UniversalPageData{
				LoggedIn:      loggedIn,
				LoggedInClass: loggedInClass,
				IsAdult:       true,
				CsrfToken:     csrf.TemplateField(r),
			},
			SearchType:                   searchType[0],
			WritingType:                  writingType[0],
			DatePosted:                   datePosted[0],
			Topics:                       topics,
			Tags:                         tags,
			WritingSearchTypes:           lib.WritingSearchTypes,
			AllWritingTopics:             allWritingTopics,
			FictionTopics:                lib.WritingTopics.Fiction,
			EssayAndBlogTopics:           lib.WritingTopics.EssaysAndBlogs,
			MemoirAndAutobiographyTopics: lib.WritingTopics.MemoirAndAutobiography,
			PoetryTopics:                 lib.WritingTopics.Poetry,
		}

		err := searchTmpl.ExecuteTemplate(w, "search_page", pageData)
		if err != nil {
			logger.Log(err)
		}
	}
}
