package libs

import "github.com/gin-gonic/gin"


func MakeResponse(c *gin.Context, statusCode int, message string, data gin.H) {
	c.JSON(statusCode, gin.H{
		"message": message,
		"data": data,
	})
}


func MakeErrorResponse(c *gin.Context, statusCode int, message string, errors ...interface{}) {
	c.JSON(statusCode, gin.H{
		"message": message,
		"errors": errors,
	})
}
