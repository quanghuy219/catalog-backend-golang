package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quanghuy219/catalog-backend-golang/controllers"
)

func Route(r *gin.Engine) {
	healthCheck := new(controllers.HealthCheck)
	r.GET("/ping", healthCheck.Ping)
}
