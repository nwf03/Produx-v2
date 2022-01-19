package db

import "gorm.io/gorm"
type Suggestion struct {
	gorm.Model
	ProductID   uint   `json:"productID"`
	Product	 Product `json:"product"`
	User        User   `json:"user"`
	UserID      uint   `json:"userID"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
type Bug struct {
	gorm.Model
	ProductID   uint   `json:"productID"`
	Product	 Product `json:"product"`
	User        User   `json:"user"`
	UserID      uint   `json:"userID"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
type Changelog struct {
	gorm.Model
	ProductID   uint   `json:"productID"`
	Product	 Product `json:"product"`
	User        User   `json:"user"`
	UserID      uint   `json:"userID"`
	Version     string `json:"version"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
type Announcement struct {
	gorm.Model
	ProductID   uint   `json:"productID"`
	Product	 Product `json:"product"`
	User        User   `json:"user"`
	UserID      uint   `json:"userID"`
	Version     string `json:"version"`
	Title       string `json:"title"`
	Description string `json:"description"`
}