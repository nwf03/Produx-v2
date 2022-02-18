package db

import (
	"errors"
	"fmt"

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
	UserLikes     []User              `json:"userLikes" gorm:"many2many:liked_suggestions;"`
	UserDislikes  []User              `json:"userDislikes" gorm:"many2many:disliked_suggestions;"`
	Comments      []SuggestionComment `json:"comments"`
	CommentsCount int64               `json:"commentsCount" gorm:"default:0"`
}
type Bug struct {
	gorm.Model
	ProductID     uint         `json:"productID"`
	Product       Product      `json:"product"`
	User          User         `json:"user"`
	UserID        uint         `json:"userID"`
	Title         string       `json:"title"`
	Description   string       `json:"description"`
	Likes         int64        `json:"likes" gorm:"default:0"`
	Dislikes      int64        `json:"dislikes" gorm:"default:0"`
	UserLikes     []User       `json:"userLikes" gorm:"many2many:liked_bugs;"`
	UserDislikes  []User       `json:"userDislikes" gorm:"many2many:disliked_bugs;"`
	Comments      []BugComment `json:"comments"`
	CommentsCount int64        `json:"commentsCount" gorm:"default:0"`
}
type Changelog struct {
	gorm.Model
	ProductID    uint    `json:"productID"`
	Product      Product `json:"product"`
	User         User    `json:"user"`
	UserID       uint    `json:"userID"`
	Version      string  `json:"version"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	Likes        int64   `json:"likes" gorm:"default:0"`
	Dislikes     int64   `json:"dislikes" gorm:"default:0"`
	UserLikes    []User  `json:"userLikes" gorm:"many2many:liked_changelogs;"`
	UserDislikes []User  `json:"userDislikes" gorm:"many2many:disliked_changelogs;"`
}
type Announcement struct {
	gorm.Model
	ProductID     uint                  `json:"productID"`
	Product       Product               `json:"product"`
	User          User                  `json:"user"`
	UserID        uint                  `json:"userID"`
	Version       string                `json:"version"`
	Title         string                `json:"title"`
	Description   string                `json:"description"`
	Likes         int64                 `json:"likes" gorm:"default:0"`
	Dislikes      int64                 `json:"dislikes" gorm:"default:0"`
	UserLikes     []User                `json:"userLikes" gorm:"many2many:liked_announcements;"`
	UserDislikes  []User                `json:"userDislikes" gorm:"many2many:disliked_announcements;"`
	Comments      []AnnouncementComment `json:"comments"`
	CommentsCount int64                 `json:"commentsCount" gorm:"default:0"`
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
type AnnouncementComment struct {
	gorm.Model
	AnnouncementID uint   `json:"announcementId"`
	User           User   `json:"user"`
	UserID         uint   `json:"userID"`
	Comment        string `json:"comment"`
}
type Post interface {
	GetID() uint
	Update(title, description string)
	AddComment(comment string, userID uint)
}
type IDGetter interface {
	GetID() uint
}
type PostLiker interface {
	Like(user User) error
	Dislike(user User) error
	RemoveDislike(user User) error
	RemoveLike(user User) error
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

func (chlog *Changelog) Update(title, description, version string) {
	if title != "" {
		chlog.Title = title
	}
	if description != "" {
		chlog.Description = description
	}
	if version != "" {
		chlog.Version = version
	}
	DB.Save(chlog)
}

func (anc *Announcement) Delete() {
	DB.Delete(&anc)
}

func (sug *Suggestion) Delete() {
	DB.Delete(&sug)
}

func (bug *Bug) Delete() {
	DB.Delete(&bug)
}

func (chlog *Changelog) Delete() {
	DB.Delete(&chlog)
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

func (chlog *Changelog) GetID() uint {
	return chlog.ID
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

func (bug *Bug) Like(user User) error {
	err := bug.RemoveDislike(user)
	if err != nil {
		return err
	}
	err = DB.Model(&bug).Association("UserLikes").Append(&user)
	DB.Table("liked_bugs").Where("bug_id = ? AND user_id = ?", bug.ID, user.ID).Count(&bug.Likes)
	DB.Save(&bug)
	if err != nil {
		return err
	}
	return nil
}
func (bug *Bug) Dislike(user User) error {
	err := bug.RemoveLike(user)
	if err != nil {
		return err
	}
	err = DB.Model(&bug).Association("UserDislikes").Append(&user)
	DB.Table("disliked_bugs").Where("bug_id = ? AND user_id = ?", bug.ID, user.ID).Count(&bug.Dislikes)
	DB.Save(&bug)
	if err != nil {
		return err
	}
	return nil
}
func (bug *Bug) RemoveLike(user User) error {
	err := DB.Model(&bug).Association("UserLikes").Delete(&user)
	DB.Table("liked_bugs").Where("bug_id = ? AND user_id = ?", bug.ID, user.ID).Count(&bug.Likes)
	DB.Save(&bug)
	if err != nil {
		return err
	}
	return nil
}

func (bug *Bug) RemoveDislike(user User) error {
	err := DB.Model(&bug).Association("UserDislikes").Delete(&user)
	DB.Table("disliked_bugs").Where("bug_id = ? AND user_id = ?", bug.ID, user.ID).Count(&bug.Dislikes)
	DB.Save(&bug)
	if err != nil {
		return err
	}
	return nil
}

func (sug *Suggestion) Like(user User) error {
	err := sug.RemoveDislike(user)
	if err != nil {
		return err
	}
	err = DB.Model(&sug).Association("UserLikes").Append(&user)
	DB.Table("liked_suggestions").Where("suggestion_id = ? AND user_id = ?", sug.ID, user.ID).Count(&sug.Likes)
	if err != nil {
		return err
	}
	return nil
}
func (sug *Suggestion) Dislike(user User) error {
	err := sug.RemoveLike(user)
	if err != nil {
		return err
	}
	err = DB.Model(&sug).Association("UserDislikes").Append(&user)
	DB.Table("disliked_suggestions").Where("suggestion_id = ? AND user_id = ?", sug.ID, user.ID).Count(&sug.Dislikes)
	if err != nil {
		return err
	}
	return nil
}
func (sug *Suggestion) RemoveLike(user User) error {
	err := DB.Model(&sug).Association("UserLikes").Delete(&user)
	DB.Table("liked_suggestions").Where("suggestion_id = ? AND user_id = ?", sug.ID, user.ID).Count(&sug.Likes)
	if err != nil {
		return err
	}
	return nil
}

func (sug *Suggestion) RemoveDislike(user User) error {
	err := DB.Model(&sug).Association("UserDislikes").Delete(&user)
	DB.Table("disliked_suggestions").Where("suggestion_id = ? AND user_id = ?", sug.ID, user.ID).Count(&sug.Dislikes)
	if err != nil {
		return err
	}
	return nil
}

func (anc *Announcement) Like(user User) error {
	err := anc.RemoveDislike(user)
	if err != nil {
		return err
	}
	err = DB.Model(&anc).Association("UserLikes").Append(&user)
	DB.Table("liked_announcements").Where("announcement_id = ? AND user_id = ?", anc.ID, user.ID).Count(&anc.Likes)
	if err != nil {
		return err
	}
	return nil
}
func (anc *Announcement) Dislike(user User) error {
	err := anc.RemoveLike(user)
	if err != nil {
		return err
	}
	err = DB.Model(&anc).Association("UserDislikes").Append(&user)
	DB.Table("disliked_announcements").Where("announcement_id = ? AND user_id = ?", anc.ID, user.ID).Count(&anc.Dislikes)
	if err != nil {
		return err
	}
	return nil
}
func (anc *Announcement) RemoveLike(user User) error {
	err := DB.Model(&anc).Association("UserLikes").Delete(&user)
	DB.Table("liked_announcements").Where("announcement_id = ? AND user_id = ?", anc.ID, user.ID).Count(&anc.Likes)
	if err != nil {
		return err
	}
	return nil
}

func (anc *Announcement) RemoveDislike(user User) error {
	err := DB.Model(&anc).Association("UserDislikes").Delete(&user)
	DB.Table("disliked_announcements").Where("announcement_id = ? AND user_id = ?", anc.ID, user.ID).Count(&anc.Dislikes)
	if err != nil {
		return err
	}
	return nil
}

func (chlog *Changelog) Like(user User) error {
	err := chlog.RemoveDislike(user)
	if err != nil {
		return err
	}
	err = DB.Model(&chlog).Association("UserLikes").Append(&user)
	DB.Table("liked_changelogs").Where("changelog_id = ? AND user_id = ?", chlog.ID, user.ID).Count(&chlog.Likes)
	if err != nil {
		return err
	}
	return nil
}
func (chlog *Changelog) Dislike(user User) error {
	err := chlog.RemoveLike(user)
	if err != nil {
		return err
	}
	err = DB.Model(&chlog).Association("UserDislikes").Append(&user)
	DB.Table("disliked_changelogs").Where("changelog_id = ? AND user_id = ?", chlog.ID, user.ID).Count(&chlog.Dislikes)
	if err != nil {
		return err
	}
	return nil
}
func (chlog *Changelog) RemoveLike(user User) error {
	err := DB.Model(&chlog).Association("UserLikes").Delete(&user)
	DB.Table("liked_changelogs").Where("changelog_id = ? AND user_id = ?", chlog.ID, user.ID).Count(&chlog.Likes)
	if err != nil {
		return err
	}
	return nil
}

func (chlog *Changelog) RemoveDislike(user User) error {
	err := DB.Model(&chlog).Association("UserDislikes").Delete(&user)
	DB.Table("disliked_changelogs").Where("changelog_id = ? AND user_id = ?", chlog.ID, user.ID).Count(&chlog.Dislikes)
	if err != nil {
		return err
	}
	return nil
}
