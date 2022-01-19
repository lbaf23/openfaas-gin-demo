package main

import (
	"demo/models"
	"demo/routers"
	"demo/setting"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
	models.Setup()
}

func main() {
	r := routers.InitRouter()
	corsMiddleware := cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	})

	gin.SetMode(setting.Config.RunMode)

	r.Use(corsMiddleware)
	r.Run(":8000")
}
