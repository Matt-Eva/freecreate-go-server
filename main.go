package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	// allowedOrigin := os.Getenv("CLIENT_ORIGIN")
	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{allowedOrigin},
	// 	AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
	// 	AllowCredentials: true,
	// 	MaxAge:           1 * time.Hour,
	// }))

	router.GET("/session", func(c *gin.Context) {
		session := sessions.Default(c)

		sessionId := session.Get("sessionId")

		fmt.Println(sessionId)

	})

	router.POST("/sessions", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("sessionId", "myId")
		session.Save()

		c.JSON(201, gin.H{"sessionId": "myId"})
	})

	router.DELETE("/sessions", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Delete("sessionId")
		session.Save()

		c.JSON(200, gin.H{"sessionId": false})
	})

	router.Run()
}
