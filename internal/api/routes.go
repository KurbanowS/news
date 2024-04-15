package api

import (
	"github.com/KurbanowS/news/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes(routes *gin.Engine) {
	if config.Conf.AppEnvIsProd {
		gin.SetMode("release")
	}
	routes.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Conten-Length", "Content-Type"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
	}))

	api := routes.Group("/api")
	{
		NewsRoutes(api)
	}

}
