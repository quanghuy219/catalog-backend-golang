package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quanghuy219/catalog-backend-golang/libs"
	"github.com/quanghuy219/catalog-backend-golang/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthenticationController struct{}

type Login struct {
	Email string `json:"Email" validate:"email" binding:"required"`
	Password string `json:"Password" binding:"required"`
}

func (a AuthenticationController) Login(c *gin.Context) {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)

	var user models.User
	if result := db.Where(&models.User{Email: json.Email}).First(&user); result != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email or password is incorrect"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(json.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email or password is incorrect"})
		return
	}

	token := libs.JwtEncode(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successfully",
		"data": gin.H{
			"id": user.ID,
			"name": user.Name,
		},
		"token": token,
	})
}

type Signup struct {
	Email string `json:"Email" validate:"email" binding:"required"`
	Password string `json:"Password" validate:"regexp=^[A-Za-z0-9]{8,}$" binding:"required"`
	ConfirmedPassword string `json:"confirmed_password" binding:"required"`
	Name string `json:"Name" binding:"required"`
}

func (a AuthenticationController) Signup(c *gin.Context) {
	var json Signup
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := c.MustGet("db").(*gorm.DB)

	var user models.User
	result := db.Where(&models.User{Email: json.Email}).First(&user)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is already being used"})
		return
	}

	if json.Password != json.ConfirmedPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password and confirmed password do not match"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 10)
	user = models.User{
		Email: json.Email,
		Name: json.Name,
		PasswordHash: string(hashedPassword),
	}
	db.Create(&user)
	token := libs.JwtEncode(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "Signup successfully",
		"data": gin.H{
			"id": user.ID,
			"name": user.Name,
		},
		"token": token,
	})
}
