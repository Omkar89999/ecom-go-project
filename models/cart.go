package models

import "time"

type Cart struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	ProductID uint      `json:"product_id"`
	Quantity  int       `json:"quantity"`

	Product   *Product  `json:"product" gorm:"foreignKey:ProductID;references:ID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Cart) TableName() string {
	return "carts"
}
