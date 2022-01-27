package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/quanghuy219/catalog-backend-golang/db"
	"github.com/quanghuy219/catalog-backend-golang/router"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

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
	r.Use(CORSMiddleware())

	// Init database connection
	db.Init()

	router.Route(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
