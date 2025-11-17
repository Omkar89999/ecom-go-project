package services

import (
	"ecom-go/config"
	"ecom-go/models"
	"errors"
)

type CartService struct{}

func (cs CartService) AddToCart(userID uint, productID uint, qty int) (models.Cart, error) {
	if qty <= 0 {
		qty = 1
	}

	// check product exist

	var product models.Product

	if err := config.DB.First(&product, productID).Error; err != nil {
		return models.Cart{}, errors.New("product not found")
	}

	// check if cart already has product
	var cart models.Cart

	err := config.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&cart).Error

	if err == nil {
		cart.Quantity += qty

		config.DB.Save(&cart)
		return cart, nil
	}

	// create new cart item

	newCart := models.Cart{
		UserID:    userID,
		ProductID: productID,
		Quantity:  qty,
	}

	if err := config.DB.Create(&newCart).Error; err != nil {
		return models.Cart{}, err

	}

	config.DB.Preload("Product").Preload("Product.Category").First(&newCart, newCart.ID)

	return newCart, nil
}

func (cs CartService) GetCart(userID uint) ([]models.Cart, error) {

	var cart []models.Cart
	if err := config.DB.Preload("Product.Category").
		Where("user_id = ?", userID).Find(&cart).Error; err != nil {
		return nil, err

	}
	return cart, nil

}

func (cs CartService) RemoveFromCart(cartID uint) error {
	if err := config.DB.Delete(&models.Cart{}, cartID).Error; err != nil {
		return err
	}
	return nil
}
