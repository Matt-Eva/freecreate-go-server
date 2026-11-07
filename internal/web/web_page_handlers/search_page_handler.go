package web_page_handlers

import (
	"fmt"
	"freecreate/internal/config"
	"freecreate/internal/lib"
	"freecreate/internal/lib/logger"
	"freecreate/internal/query_handlers"
	"freecreate/internal/web/web_auth"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valkey-io/valkey-go"
)

func SearchPageHandler(searchTmpl *template.Template, sessionStore *sessions.CookieStore, valkeyClient valkey.Client, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries) http.HandlerFunc {
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

		searchTypeParams := query["searchType"]
		searchType := "Both"
		if len(searchTypeParams) > 0{
			searchType = searchTypeParams[0]
		}
		fmt.Println(searchType)

		writingTypeParams := query["writingType"]
		writingType := "All"
		if len(writingTypeParams) > 0{
			writingType = writingTypeParams[0]
		}
		fmt.Println(writingType)

		datePostedParams := query["datePosted"]
		datePosted := "All Time"
		if len(datePostedParams) >0{
			datePosted = datePostedParams[0]
		}
		fmt.Println(datePosted)

		topics := query["topic"]

		fmt.Println(topics)

		tags := query["tag"]
		fmt.Println(tags)

		searchResults, searchErr := query_handlers.HandleSearch(ctx, valkeyClient, pgCore, pgCoreQueries, searchType, writingType, datePosted, topics, tags)
		if searchErr != nil {
			
		}

		type PageData struct {
			UniversalPageData
			SearchType                   string
			WritingType                  string
			DatePosted                   string
			Topics                       []string
			Tags                         []string
			SearchTypes                  []string
			WritingSearchTypes           []string
			DatePostedSearchTypes        []string
			AllWritingTopics             []string
			FictionTopics                []string
			EssayAndBlogTopics           []string
			MemoirAndAutobiographyTopics []string
			PoetryTopics                 []string
		}

		pageData := PageData{
			UniversalPageData: UniversalPageData{
				LoggedIn:      loggedIn,
				LoggedInClass: loggedInClass,
				IsAdult:       true,
				CsrfToken:     csrf.TemplateField(r),
			},
			SearchType:                   searchType,
			WritingType:                  writingType,
			DatePosted:                   datePosted,
			Topics:                       searchResults.Topics,
			Tags:                         searchResults.Tags,
			SearchTypes:                  lib.SearchTypes,
			WritingSearchTypes:           lib.WritingSearchTypes,
			DatePostedSearchTypes:        lib.DatePostedSearchTypes,
			AllWritingTopics:             lib.AllWritingTopics,
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
