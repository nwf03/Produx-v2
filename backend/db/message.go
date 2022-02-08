package db

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ProductId uint   `json:"product_id"`
	UserID    uint   `json:"user_id"`
	User      User   `json:"user"`
	Message   string `json:"message"`
}
