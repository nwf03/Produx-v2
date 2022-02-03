package db

import (
	"errors"
	"fmt"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Suggestion struct {
	gorm.Model
	ProductID     uint                `json:"productID"`
	Product       Product             `json:"product"`
	User          User                `json:"user"`
	UserID        uint                `json:"userID"`
	Title         string              `json:"title"`
	Description   string              `json:"description"`
	Likes         int64               `json:"likes" gorm:"default:0"`
	Dislikes      int64               `json:"dislikes" gorm:"default:0"`
	UserLikes     pq.Int64Array       `json:"userLikes" gorm:"type:bigint[]"`
	UserDislikes  pq.Int64Array       `json:"userDislikes" gorm:"type:bigint[]"`
	Comments      []SuggestionComment `json:"comments"`
	CommentsCount int64               `json:"commentsCount" gorm:"default:0"`
}
type Bug struct {
	gorm.Model
	ProductID     uint          `json:"productID"`
	Product       Product       `json:"product"`
	User          User          `json:"user"`
	UserID        uint          `json:"userID"`
	Title         string        `json:"title"`
	Description   string        `json:"description"`
	Likes         int64         `json:"likes" gorm:"default:0"`
	Dislikes      int64         `json:"dislikes" gorm:"default:0"`
	UserLikes     pq.Int64Array `json:"userLikes" gorm:"type:bigint[]"`
	UserDislikes  pq.Int64Array `json:"userDislikes" gorm:"type:bigint[]"`
	Comments      []BugComment  `json:"comments"`
	CommentsCount int64         `json:"commentsCount" gorm:"default:0"`
}
type Changelog struct {
	gorm.Model
	ProductID    uint          `json:"productID"`
	Product      Product       `json:"product"`
	User         User          `json:"user"`
	UserID       uint          `json:"userID"`
	Version      string        `json:"version"`
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	Likes        int64         `json:"likes" gorm:"default:0"`
	Dislikes     int64         `json:"dislikes" gorm:"default:0"`
	UserLikes    pq.Int64Array `json:"userLikes" gorm:"type:bigint[]"`
	UserDislikes pq.Int64Array `json:"userDislikes" gorm:"type:bigint[]"`
}
type Announcement struct {
	gorm.Model
	ProductID    uint          `json:"productID"`
	Product      Product       `json:"product"`
	User         User          `json:"user"`
	UserID       uint          `json:"userID"`
	Version      string        `json:"version"`
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	Likes        int64         `json:"likes" gorm:"default:0"`
	Dislikes     int64         `json:"dislikes" gorm:"default:0"`
	UserLikes    pq.Int64Array `json:"userLikes" gorm:"type:bigint[]"`
	UserDislikes pq.Int64Array `json:"userDislikes" gorm:"type:bigint[]"`
}

type SuggestionComment struct {
	gorm.Model
	SuggestionID uint   `json:"suggestionID"`
	User         User   `json:"user"`
	UserID       uint   `json:"userID"`
	Comment      string `json:"comment"`
}
type BugComment struct {
	gorm.Model
	BugID   uint   `json:"bugID"`
	User    User   `json:"user"`
	UserID  uint   `json:"userID"`
	Comment string `json:"comment"`
}
type Post interface {
	GetID() uint
	Update(title, description string)
	AddComment(comment string, userID uint)
}

func (sug *Suggestion) AddComment(comment string, userID uint) {
	var User User
	DB.First(&User, userID)
	sug.Comments = append(sug.Comments, SuggestionComment{
		Comment:      comment,
		User:         User,
		UserID:       userID,
		SuggestionID: sug.ID,
	})
	DB.Save(&sug)
}
func (bug *Bug) AddComment(comment string, userID uint) {
	var User User
	DB.First(&User, userID)
	bug.Comments = append(bug.Comments, BugComment{
		Comment: comment,
		User:    User,
		UserID:  userID,
		BugID:   bug.ID,
	})
	DB.Save(&bug)
}

type AdminPost interface {
	Update(title, description, version string)
	GetID() uint
}
type DeletePost interface {
	Delete()
	GetID() uint
}

func (anc *Announcement) Update(newTitle, newDescription, newVersion string) {
	fmt.Println("Product being updated")
	if newTitle != "" {
		anc.Title = newTitle
	}
	if newDescription != "" {
		anc.Description = newDescription
	}
	if newVersion != "" {
		anc.Version = newVersion
	}
	DB.Save(anc)
}

func (sug *Suggestion) Update(title, description string) {
	if title != "" {
		sug.Title = title
	}
	if description != "" {
		sug.Description = description
	}
	DB.Save(sug)
}

func (bug *Bug) Update(title, description string) {
	fmt.Println("Product being updated")

	if title != "" {
		bug.Title = title
	}
	if description != "" {
		bug.Description = description
	}
	DB.Save(bug)
}

func (changelog *Changelog) Update(title, description, version string) {
	if title != "" {
		changelog.Title = title
	}
	if description != "" {
		changelog.Description = description
	}
	if version != "" {
		changelog.Version = version
	}
	DB.Save(changelog)
}

func (anc *Announcement) Delete() {
	DB.Delete(anc)
}

func (sug *Suggestion) Delete() {
	DB.Delete(sug)
}

func (bug *Bug) Delete() {
	DB.Delete(bug)
}

func (changelog *Changelog) Delete() {
	DB.Delete(changelog)
}

func (anc *Announcement) GetID() uint {
	return anc.ID
}

func (sug *Suggestion) GetID() uint {
	return sug.ID
}

func (bug *Bug) GetID() uint {
	return bug.ID
}

func (changelog *Changelog) GetID() uint {
	return changelog.ID
}
func (bc *BugComment) BeforeCreate(tx *gorm.DB) error {
	var Bug Bug
	err := DB.First(&Bug, bc.BugID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("Bug not found")
	}
	Bug.CommentsCount++
	DB.Save(&Bug)
	return nil
}
func (sc *SuggestionComment) BeforeCreate(tx *gorm.DB) error {
	var Suggestion Suggestion
	err := DB.First(&Suggestion, sc.SuggestionID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("Suggestion not found")
	}
	Suggestion.CommentsCount++
	DB.Save(&Suggestion)
	return nil
}
