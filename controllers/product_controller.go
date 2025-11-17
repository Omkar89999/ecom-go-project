package controllers

import (
	"ecom-go/config"
	"ecom-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check category exists
	var category models.Category
	if err := config.DB.First(&category, product.CategoryID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Category ID"})
		return
	}

	config.DB.Create(&product)
	config.DB.Preload("Category").First(&product, product.ID)
	c.JSON(http.StatusCreated, gin.H{"product": product})
}

func GetProducts(c *gin.Context) {
	var products []models.Product
	config.DB.Preload("Category").Find(&products)

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")

	var product models.Product
	if config.DB.Preload("Category").First(&product, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var product models.Product
	if config.DB.First(&product, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var input models.Product
	c.ShouldBindJSON(&input)

	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.Stock = input.Stock
	product.CategoryID = input.CategoryID

	config.DB.Save(&product)

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Product{}, id)

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
