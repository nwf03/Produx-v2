package db

import (
	"errors"
	"fmt"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ProductID     uint           `json:"productID"`
	Product       Product        `json:"product"`
	User          User           `json:"user"`
	UserID        uint           `json:"userID"`
	Title         string         `json:"title"`
	Description   string         `json:"description"`
	Likes         int64          `json:"likes" gorm:"default:0"`
	Dislikes      int64          `json:"dislikes" gorm:"default:0"`
	UserLikes     []User         `json:"userLikes" gorm:"many2many:liked_posts;"`
	UserDislikes  []User         `json:"userDislikes" gorm:"many2many:disliked_posts;"`
	Comments      []Comment      `json:"comments"`
	CommentsCount int64          `json:"commentsCount" gorm:"default:0"`
	Type          pq.StringArray `json:"type" gorm:"type:text[]"`
	Images        pq.StringArray `json:"images" gorm:"type:text[]"`
}

func ValidType(t string) bool {
	fmt.Println("checking type t: ", t)
	fmt.Println(len(t))
	switch t {
	case "suggestions", "bugs", "announcements", "done", "working-on", "under-review", "changelogs":
		return true
	}
	return false
}

func checkTypes(t []string) bool {
	for _, s := range t {
		if !ValidType(s) {
			return false
		}
	}
	return true
}

func (p *Post) BeforeCreate(db *gorm.DB) error {
	if !checkTypes(p.Type) {
		return errors.New("invalid type")
	}
	return nil
}

func (p *Post) AddType(t string) error {
	if !ValidType(t) {
		return errors.New("invalid type")
	}
	p.Type = append(p.Type, t)
	DB.Save(&p)
	return nil
}

func (p *Post) RemoveType(t string) error {
	if !ValidType(t) {
		return errors.New("invalid type")
	}
	var newTypes pq.StringArray
	for _, field := range p.Type {
		if field != t {
			newTypes = append(newTypes, field)
		}
	}
	fmt.Println("new typesss: ", newTypes)
	p.Type = newTypes
	DB.Save(&p)
	return nil
}

type Comment struct {
	gorm.Model
	PostID  uint   `json:"postID"`
	User    User   `json:"user"`
	UserID  uint   `json:"userID"`
	Comment string `json:"comment"`
}

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

func (p *Post) Delete() {
	DB.Delete(&p)
}

func (p *Post) GetID() uint {
	return p.ID
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

type test struct {
	Name     string `json:"nwf"`
	LastName string `json:"nwf"`
}
