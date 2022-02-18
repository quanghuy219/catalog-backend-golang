package middlewares

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/quanghuy219/catalog-backend-golang/libs"
	"github.com/quanghuy219/catalog-backend-golang/models"
	"gorm.io/gorm"
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

func DatabaseMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func JwtAuthMiddlewware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			libs.MakeErrorResponse(c, http.StatusUnauthorized, "Missing authorization header")
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			libs.MakeErrorResponse(c, http.StatusBadRequest, "Invalid authorization header format")
			c.Abort()
			return
		}
		token, err := libs.ParseJwtToken(parts[1])
		if err != nil {
			libs.MakeErrorResponse(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
			user_id, err := strconv.Atoi(claims.Audience)
			if err != nil {
				libs.MakeErrorResponse(c, http.StatusBadRequest, "Invalid token")
				c.Abort()
				return
			}
			var user models.User
			db := c.MustGet("db").(*gorm.DB)
			if err := db.First(&user, user_id).Error; err != nil {
				libs.MakeErrorResponse(c, http.StatusBadRequest, "Invalid token")
				c.Abort()
				return
			}
			c.Set("user", user)
			c.Next()
		} else {
			libs.MakeErrorResponse(c, http.StatusBadRequest, "Invalid token")
			c.Abort()
			return
		}
	}
}
