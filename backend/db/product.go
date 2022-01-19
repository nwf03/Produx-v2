package db

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)
type Product struct {
	gorm.Model
	UserID uint   `json:"userID,omitempty"`
	Name   string `json:"name,omitempty"`
	Users []User `json:"users,omitempty" gorm:"many2many:followed_products;"`
	Description string         `json:"description,omitempty"`
	Suggestions []Suggestion   `json:"suggestions,omitempty"`
	Bugs        []Bug          `json:"bugs,omitempty"`
	Changelogs  []Changelog    `json:"changeLogs,omitempty"`
	Announcements []Announcement `json:"announcements,omitempty"`
	Images      pq.StringArray `json:"images" gorm:"type:text[]"`
	UserLikes []User `json:"user_likes" gorm:"many2many:likes;"`
	LikesCount int64 `json:"likes" gorm:"default:0"`
	Verified bool `json:"verified" gorm:"default:false"`
}
