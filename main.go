package main

import (
	"log"

	"github.com/KurbanowS/news/config"
	"github.com/KurbanowS/news/internal/api"
	"github.com/KurbanowS/news/internal/store"
	"github.com/KurbanowS/news/internal/store/pgx"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	defer store.Init().(*pgx.PgxStore).Close()

	routes := gin.Default()
	api.Routes(routes)
	if err := routes.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
