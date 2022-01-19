package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name"`
	Password string `json:"password"`
	Email string `json:"email"`
	Pfp string `json:"pfp"`
	Products []Product `json:"products" gorm:"foreignKey:UserID"`
	FollowedProducts []Product `json:"followed_products" gorm:"many2many:followed_products;"`
	Suggestions []Suggestion `json:"suggestions,omitempty"`
	Bugs []Bug `json:"bugs,omitempty"`
	Changelogs []Changelog `json:"changeLogs,omitempty"`
	LikedProducts []Product `json:"liked_products" gorm:"many2many:likes;"`
}