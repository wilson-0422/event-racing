package main

import (
	"event-racing/src/config"
	"event-racing/src/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	Seed()

	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/static", "static")

	store := cookie.NewStore([]byte(config.SessionSecret))
	r.Use(sessions.Sessions(config.SessionName, store))

	routes.SetupRoutes(r)

	r.Run(":" + config.Port)
}
