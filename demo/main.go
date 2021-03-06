package main

import (
	"demo/models"
	"demo/routers"
	"demo/settings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	settings.Setup()
	models.Setup()
}

func main() {
	r := routers.InitRouter(settings.Config.RunMode)
	corsMiddleware := cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	})

	gin.SetMode(settings.Config.RunMode)

	r.Use(corsMiddleware)
	r.Run(":8000")
}
