package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quanghuy219/catalog-backend-golang/controllers"
)

func Route(r *gin.Engine) {
	healthCheck := new(controllers.HealthCheck)
	authController := new(controllers.AuthenticationController)
	categoryController := new(controllers.CategoryController)

	r.GET("/ping", healthCheck.Ping)
	r.POST("/login", authController.Login)
	r.POST("/users", authController.Signup)

	r.GET("/categories", categoryController.GetAllCategories)
}
