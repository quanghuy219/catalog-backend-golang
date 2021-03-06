package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/quanghuy219/catalog-backend-golang/database"
	"github.com/quanghuy219/catalog-backend-golang/middlewares"
	"github.com/quanghuy219/catalog-backend-golang/router"
)

func main() {
	//Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
	}

	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	// Init middlewares
	r.Use(middlewares.CORSMiddleware())

	// Init database connection
	db := database.Init()
	r.Use(middlewares.DatabaseMiddleware(db))

	router.Route(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
