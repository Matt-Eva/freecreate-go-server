package main

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{}
	}))

	router.GET("/session", func(c *gin.Context){
		session := sessions.Default(c)

		sessionId := session.Get("sessionId")

		fmt.Println(sessionId)
		
	})

	router.POST("/sessions", func (c *gin.Context){})

	router.DELETE("/sessions", func(c *gin.Context){})
}