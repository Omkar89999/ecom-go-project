package controllers

import (
	"ecom-go/config"
	"ecom-go/models"
	"ecom-go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {

	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// hash  password

	hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	user := models.User{

		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashed),
	}

	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User Registered Successfully",
		"user":    user,
	})

}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {

	var input LoginInput

	c.ShouldBindJSON(&input)

	var user models.User

	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Email"})
		return
	}

	// check Password

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Password"})

		return
	}

	token, _ := utils.GenerateToken(user.ID)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successfull",
		"token":   token,
	})

}
