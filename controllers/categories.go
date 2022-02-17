package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quanghuy219/catalog-backend-golang/models"
	"gorm.io/gorm"
)

type CategoryController struct{}

func (ctrl CategoryController) GetAllCategories(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var categories []models.Category
	db.Find(&categories)
	c.JSON(http.StatusOK, gin.H{
		"message": "Get all available categories",
		"data": gin.H{
			"categories": categories,
		},
	})
}

type CategorySchema struct {
	Name string `json:"name" binding:"required" validate:"min:1,max:250"`
}

func (ctrl CategoryController) CreateCategory(c *gin.Context) {
	var json CategorySchema
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := c.MustGet("db").(*gorm.DB)
	category := models.Category{
		Name: json.Name,
	}
	db.Create(&category)
	c.JSON(http.StatusOK, gin.H{
		"message": "New category created",
		"data": gin.H{
			"category": category,
		},
	})
}
