package routes

import (
	"encoding/json"
	"fmt"
	"freecreate/logger"
	"freecreate/middleware"
	"freecreate/web_api_handlers"
	"freecreate/web_page_handlers"
	"html/template"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/resend/resend-go/v2"
	"github.com/valkey-io/valkey-go"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"gorm.io/gorm"
)



func CreateRouter(sessionStore *sessions.CookieStore, gormPGClient *gorm.DB, mongoClient *mongo.Client, valkeyClient valkey.Client, resendClient *resend.Client) *chi.Mux {
	router := chi.NewRouter()

	environment := os.Getenv("ENVIRONMENT")
	csrfKey := os.Getenv("CSRF_KEY")	
	var csrfMiddleware func(http.Handler) http.Handler
	if environment == "DEVELOPMENT"{
		fmt.Println("DEVELOPMENT")
		csrfMiddleware = csrf.Protect([]byte(csrfKey), csrf.Secure(false), csrf.TrustedOrigins([]string{"localhost:8080"}))
	} else {
		csrfMiddleware = csrf.Protect([]byte(csrfKey))
	}
	
	router.Use(csrfMiddleware)

	fileServer := http.FileServer(http.Dir("static"))
 	cachedFileServer := middleware.CacheControlHandler(fileServer)
	
	router.Handle("/static/*", http.StripPrefix("/static/", cachedFileServer))

	router.Get("/get-csrf", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-CSRF-Token", csrf.Token(r))
		w.Header().Set("Access-Control-Expose-Headers", "X-CSRF-Token")
	})

	templates := template.Must(template.ParseGlob("templates/*html"));
	
	// testTmpl := template.Must(template.ParseFiles("templates/pages/test/test.html", "templates/components/globals.html", "templates/components/header.html"))
	router.Get("/test", web_page_handlers.TestHandler(templates))
	router.Post("/test", web_page_handlers.TestHandler(templates))

	// homeTmpl := template.Must(template.ParseFiles("templates/pages/home/home.html", "templates/pages/home/searchBox.html", "templates/pages/home/contentCard.html", "templates/components/header.html", "templates/components/globals.html"))
	router.Get("/", web_page_handlers.HomeHandler(templates))

	// aboutTmpl := template.Must(template.ParseFiles("templates/pages/about/about.html", "templates/components/header.html", "templates/components/globals.html"))
	// router.Get("/about", web_page_handlers.AboutHandler(aboutTmpl))

	// loginTmpl := template.Must(template.ParseFiles("templates/pages/login/login.html", "templates/components/globals.html", "templates/components/header.html"))
	// router.Get("/login", web_page_handlers.LoginHandler(loginTmpl))

	// profileTmpl := template.Must(template.ParseFiles("templates/pages/profile/profile.html", "templates/components/header.html", "templates/components/globals.html"))
	// router.Get("/profile", web_page_handlers.ProfileHandler(profileTmpl))

	// donateTmpl := template.Must(template.ParseFiles("templates/pages/donate/donate.html", "templates/components/header.html", "templates/components/globals.html"))
	// router.Get("/donate", web_page_handlers.DonateHandler(donateTmpl))

	router.Route("/web-api", func(r chi.Router) {

		router.Post("/login", web_api_handlers.LoginHandler(sessionStore, gormPGClient))

		router.Post("/signup", web_api_handlers.SignupHandler(sessionStore, gormPGClient))

		router.Delete("/logout", web_api_handlers.LogoutHandler(sessionStore, gormPGClient))

		router.Get("/reauth", web_api_handlers.ReAuthHandler(sessionStore, gormPGClient))

		router.Delete("/delete-account", web_api_handlers.DeleteAccountHandler(sessionStore, gormPGClient))

		router.Post("/creator", web_api_handlers.CreateCreatorHandler(sessionStore, gormPGClient))

		router.Delete("/creator/{creatorId}", web_api_handlers.DeleteCreatorHandler(sessionStore, gormPGClient))

		router.Get("/user-creators", web_api_handlers.GetUserCreatorHandlers(sessionStore, gormPGClient))

		router.Post("/writing", web_api_handlers.CreateWritingHandler(sessionStore, gormPGClient))

		router.Patch("/writing", web_api_handlers.UpdateWritingHandler(sessionStore, gormPGClient))

		router.Get("/my-writing", web_api_handlers.GetMyWritingHandler(sessionStore, gormPGClient))

		router.Get("/edit-writing/{writingUUID}", web_api_handlers.GetEditWritingHandler(sessionStore, gormPGClient))

		// router.Post("/createOTP", handlers.CreateOTPHandler(resendClient, valkeyClient))

		// router.Post("/email", handlers.EmailHandler(resendClient))

		router.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("hit route hello")

			type Response struct {
				Message string `json:"message"`
			}

			response := Response{
				Message: "Hello world!",
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)
		})

		router.Post("/hello", func(w http.ResponseWriter, r *http.Request) {
			type Body struct {
				Message string `json:"message"`
			}

			var body Body

			jErr := json.NewDecoder(r.Body).Decode(&body)
			if jErr != nil {
				logger.Log(jErr)
				http.Error(w, jErr.Error(), http.StatusUnprocessableEntity)
				return
			}

			fmt.Println(body.Message)

			type Response struct {
				Message string `json:"message"`
			}

			response := Response{
				Message: body.Message,
			}

			res, mErr := json.Marshal(response)
			if mErr != nil {
				logger.Log(mErr)
				http.Error(w, mErr.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write(res)
		})

	})

	return router
}

// router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("hit route")

// 		type Response struct {
// 			Message string `json:"message"`
// 		}

// 		response := Response{
// 			Message: "Hello world!",
// 		}

// 		w.WriteHeader(http.StatusOK)
// 		json.NewEncoder(w).Encode(response)
// 	}).Methods("GET")

// 	router.HandleFunc("/session", func(w http.ResponseWriter, r *http.Request) {

// 	}).Methods("GET")

// 	router.HandleFunc("/session", func(w http.ResponseWriter, r *http.Request) {

// 	}).Methods("POST")

// 	router.HandleFunc("/session", func(w http.ResponseWriter, r *http.Request) {

// 	}).Methods("DELETE")
