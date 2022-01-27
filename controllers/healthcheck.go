package controllers

import "github.com/gin-gonic/gin"

type HealthCheck struct{}

func (ctr HealthCheck) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
