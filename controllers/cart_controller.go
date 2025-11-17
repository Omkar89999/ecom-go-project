package controllers

import (
	"ecom-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var CartService = services.CartService{}

type AddCartInput struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

func AddToCart(c *gin.Context) {

	userID := c.GetUint("user_id")

	var input AddCartInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cart, err := CartService.AddToCart(userID, input.ProductID, input.Quantity)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Added to cart",
		"cart":    cart,
	})

}

func ViewCart(c *gin.Context) {
	UserID := c.GetUint("user_id")

	cart, err := CartService.GetCart(UserID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cart": cart})
}

func RemoveFromCart(c *gin.Context) {

	idStr := c.Param("id")

	cartID, _ := strconv.Atoi(idStr)

	err := CartService.RemoveFromCart(uint(cartID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to remove item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item removed"})
}
