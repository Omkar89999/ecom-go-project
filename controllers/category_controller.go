package controllers

import (
	"ecom-go/config"
	"ecom-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {

	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category already exists"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"category": category})
}

func GetCategories(c *gin.Context) {
	var categories []models.Category

	config.DB.Find(&categories)

	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

func GetCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category

	if err := config.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"category": category})
}

func UpdateCategory(c *gin.Context) {

	id := c.Param("id")

	var category models.Category

	if config.DB.First(&category, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}

	var input models.Category
	c.ShouldBindJSON(&input)

	category.Name = input.Name

	config.DB.Save(&category)

	c.JSON(http.StatusOK, gin.H{"category": category})
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	config.DB.Delete(&models.Category{}, id)

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}
