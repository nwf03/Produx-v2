package db

import (
	"errors"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ProductID     uint      `json:"productID"`
	Product       Product   `json:"product"`
	User          User      `json:"user"`
	UserID        uint      `json:"userID"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Likes         int64     `json:"likes" gorm:"default:0"`
	Dislikes      int64     `json:"dislikes" gorm:"default:0"`
	UserLikes     []User    `json:"userLikes" gorm:"many2many:liked_suggestions;"`
	UserDislikes  []User    `json:"userDislikes" gorm:"many2many:disliked_suggestions;"`
	Comments      []Comment `json:"comments"`
	CommentsCount int64     `json:"commentsCount" gorm:"default:0"`
	Type          string    `json:"type"`
}

func ValidType(t string) bool {
	switch t {
	case "suggestion", "bug", "announcement":
		return true
	}
	return false
}

func (p *Post) BeforeCreate(db *gorm.DB) error {
	if !ValidType(p.Type) {
		return errors.New("type is required")
	}
	return nil
}

//type Bug struct {
//	gorm.Model
//	ProductID     uint         `json:"productID"`
//	Product       Product      `json:"product"`
//	User          User         `json:"user"`
//	UserID        uint         `json:"userID"`
//	Title         string       `json:"title"`
//	Description   string       `json:"description"`
//	Likes         int64        `json:"likes" gorm:"default:0"`
//	Dislikes      int64        `json:"dislikes" gorm:"default:0"`
//	UserLikes     []User       `json:"userLikes" gorm:"many2many:liked_bugs;"`
//	UserDislikes  []User       `json:"userDislikes" gorm:"many2many:disliked_bugs;"`
//	Comments      []BugComment `json:"comments"`
//	CommentsCount int64        `json:"commentsCount" gorm:"default:0"`
//}
//type Announcement struct {
//	gorm.Model
//	ProductID     uint                  `json:"productID"`
//	Product       Product               `json:"product"`
//	User          User                  `json:"user"`
//	UserID        uint                  `json:"userID"`
//	Version       string                `json:"version"`
//	Title         string                `json:"title"`
//	Description   string                `json:"description"`
//	Likes         int64                 `json:"likes" gorm:"default:0"`
//	Dislikes      int64                 `json:"dislikes" gorm:"default:0"`
//	UserLikes     []User                `json:"userLikes" gorm:"many2many:liked_announcements;"`
//	UserDislikes  []User                `json:"userDislikes" gorm:"many2many:disliked_announcements;"`
//	Comments      []AnnouncementComment `json:"comments"`
//	CommentsCount int64                 `json:"commentsCount" gorm:"default:0"`
//}
//
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
type Comment struct {
	gorm.Model
	PostID  uint   `json:"suggestionID"`
	User    User   `json:"user"`
	UserID  uint   `json:"userID"`
	Comment string `json:"comment"`
}

//type SuggestionComment struct {
//	gorm.Model
//	SuggestionID uint   `json:"suggestionID"`
//	User         User   `json:"user"`
//	UserID       uint   `json:"userID"`
//	Comment      string `json:"comment"`
//}
//type BugComment struct {
//	gorm.Model
//	BugID   uint   `json:"bugID"`
//	User    User   `json:"user"`
//	UserID  uint   `json:"userID"`
//	Comment string `json:"comment"`
//}
//type AnnouncementComment struct {
//	gorm.Model
//	AnnouncementID uint   `json:"announcementId"`
//	User           User   `json:"user"`
//	UserID         uint   `json:"userID"`
//	Comment        string `json:"comment"`
//}
type Poster interface {
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

func (p *Post) AddComment(comment string, userID uint) {
	var User User
	DB.First(&User, userID)
	p.Comments = append(p.Comments, Comment{
		Comment: comment,
		User:    User,
		UserID:  userID,
		PostID:  p.ID,
	})
	DB.Save(&p)
}

// func (bug *Bug) AddComment(comment string, userID uint) {
// 	var User User
// 	DB.First(&User, userID)
// 	bug.Comments = append(bug.Comments, BugComment{
// 		Comment: comment,
// 		User:    User,
// 		UserID:  userID,
// 		BugID:   bug.ID,
// 	})
// 	DB.Save(&bug)
// }
//
type AdminPost interface {
	Update(title, description, version string)
	GetID() uint
}
type DeletePost interface {
	Delete()
	GetID() uint
}

func (p *Post) Update(newTitle, newDescription string) {
	if newTitle != "" {
		p.Title = newTitle
	}
	if newDescription != "" {
		p.Description = newDescription
	}
	DB.Save(&p)
}

// func (sug *Suggestion) Update(title, description string) {
// 	if title != "" {
// 		sug.Title = title
// 	}
// 	if description != "" {
// 		sug.Description = description
// 	}
// 	DB.Save(sug)
// }
//
// func (bug *Bug) Update(title, description string) {
// 	fmt.Println("Product being updated")
//
// 	if title != "" {
// 		bug.Title = title
// 	}
// 	if description != "" {
// 		bug.Description = description
// 	}
// 	DB.Save(bug)
// }

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
	DB.Save(&chlog)
}

func (p *Post) Delete() {
	DB.Delete(&p)
}

// func (anc *Announcement) Delete() {
// 	DB.Delete(&anc)
// }
//
// func (sug *Suggestion) Delete() {
// 	DB.Delete(&sug)
// }
//
// func (bug *Bug) Delete() {
// 	DB.Delete(&bug)
// }

func (chlog *Changelog) Delete() {
	DB.Delete(&chlog)
}
func (p *Post) GetID() uint {
	return p.ID
}

//
// func (anc *Announcement) GetID() uint {
// 	return anc.ID
// }
//
// func (sug *Suggestion) GetID() uint {
// 	return sug.ID
// }
//
// func (bug *Bug) GetID() uint {
// 	return bug.ID
// }
//
func (chlog *Changelog) GetID() uint {
	return chlog.ID
}
func (c *Comment) BeforeCreate(tx *gorm.DB) error {
	var post Post
	err := DB.First(&post, c.PostID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("post not found")
	}
	post.CommentsCount++
	DB.Save(&post)
	return nil
}

// func (sc *SuggestionComment) BeforeCreate(tx *gorm.DB) error {
// 	var Suggestion Suggestion
// 	err := DB.First(&Suggestion, sc.SuggestionID).Error
// 	if errors.Is(err, gorm.ErrRecordNotFound) {
// 		return errors.New("Suggestion not found")
// 	}
// 	Suggestion.CommentsCount++
// 	DB.Save(&Suggestion)
// 	return nil
// }

func (p *Post) Like(user User) error {
	err := p.RemoveDislike(user)
	if err != nil {
		return err
	}
	err = DB.Model(&p).Association("UserLikes").Append(&user)
	DB.Table("liked_posts").Where("post_id = ? AND user_id = ?", p.ID, user.ID).Count(&p.Likes)
	DB.Save(&p)
	if err != nil {
		return err
	}
	return nil
}
func (p *Post) Dislike(user User) error {
	err := p.RemoveLike(user)
	if err != nil {
		return err
	}
	err = DB.Model(&p).Association("UserDislikes").Append(&user)
	DB.Table("disliked_posts").Where("post_id = ? AND user_id = ?", p.ID, user.ID).Count(&p.Dislikes)
	DB.Save(&p)
	if err != nil {
		return err
	}
	return nil
}
func (p *Post) RemoveLike(user User) error {
	err := DB.Model(&p).Association("UserLikes").Delete(&user)
	DB.Table("liked_posts").Where("post_id = ? AND user_id = ?", p.ID, user.ID).Count(&p.Likes)
	DB.Save(&p)
	if err != nil {
		return err
	}
	return nil
}

func (p *Post) RemoveDislike(user User) error {
	err := DB.Model(&p).Association("UserDislikes").Delete(&user)
	DB.Table("disliked_posts").Where("post_id = ? AND user_id = ?", p.ID, user.ID).Count(&p.Dislikes)
	DB.Save(&p)
	if err != nil {
		return err
	}
	return nil
}

//func (sug *Suggestion) Like(user User) error {
//	err := sug.RemoveDislike(user)
//	if err != nil {
//		return err
//	}
//	err = DB.Model(&sug).Association("UserLikes").Append(&user)
//	DB.Table("liked_suggestions").Where("suggestion_id = ? AND user_id = ?", sug.ID, user.ID).Count(&sug.Likes)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//func (sug *Suggestion) Dislike(user User) error {
//	err := sug.RemoveLike(user)
//	if err != nil {
//		return err
//	}
//	err = DB.Model(&sug).Association("UserDislikes").Append(&user)
//	DB.Table("disliked_suggestions").Where("suggestion_id = ? AND user_id = ?", sug.ID, user.ID).Count(&sug.Dislikes)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//func (sug *Suggestion) RemoveLike(user User) error {
//	err := DB.Model(&sug).Association("UserLikes").Delete(&user)
//	DB.Table("liked_suggestions").Where("suggestion_id = ? AND user_id = ?", sug.ID, user.ID).Count(&sug.Likes)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (sug *Suggestion) RemoveDislike(user User) error {
//	err := DB.Model(&sug).Association("UserDislikes").Delete(&user)
//	DB.Table("disliked_suggestions").Where("suggestion_id = ? AND user_id = ?", sug.ID, user.ID).Count(&sug.Dislikes)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (anc *Announcement) Like(user User) error {
//	err := anc.RemoveDislike(user)
//	if err != nil {
//		return err
//	}
//	err = DB.Model(&anc).Association("UserLikes").Append(&user)
//	DB.Table("liked_announcements").Where("announcement_id = ? AND user_id = ?", anc.ID, user.ID).Count(&anc.Likes)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//func (anc *Announcement) Dislike(user User) error {
//	err := anc.RemoveLike(user)
//	if err != nil {
//		return err
//	}
//	err = DB.Model(&anc).Association("UserDislikes").Append(&user)
//	DB.Table("disliked_announcements").Where("announcement_id = ? AND user_id = ?", anc.ID, user.ID).Count(&anc.Dislikes)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//func (anc *Announcement) RemoveLike(user User) error {
//	err := DB.Model(&anc).Association("UserLikes").Delete(&user)
//	DB.Table("liked_announcements").Where("announcement_id = ? AND user_id = ?", anc.ID, user.ID).Count(&anc.Likes)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (anc *Announcement) RemoveDislike(user User) error {
//	err := DB.Model(&anc).Association("UserDislikes").Delete(&user)
//	DB.Table("disliked_announcements").Where("announcement_id = ? AND user_id = ?", anc.ID, user.ID).Count(&anc.Dislikes)
//	if err != nil {
//		return err
//	}
//	return nil
//}

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
