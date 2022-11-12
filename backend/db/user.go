package db

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name             string    `json:"name"`
	Password         string    `json:"-"`
	Email            string    `json:"email"`
	Pfp              string    `json:"pfp"`
	Products         []Product `json:"products" gorm:"foreignKey:UserID"`
	FollowedProducts []Product `json:"followed_products" gorm:"many2many:followed_products;"`
	Posts            []Post    `json:"posts,omitempty"`
	LikedProducts    []Product `json:"liked_products" gorm:"many2many:likes;"`

	LikedPosts    []Post    `json:"liked_posts" gorm:"many2many:liked_posts;"`
	DislikedPosts []Post    `json:"disliked_posts" gorm:"many2many:disliked_posts;"`
	Messages      []Message `json:"messages,omitempty"`
}

func (user *User) LikePost(post *Post) {
	user.RemoveDislike(post)
	user.LikedPosts = append(user.LikedPosts, *post)
	DB.Save(&user)
}

func (user *User) RemoveLike(post *Post) {
	err := DB.Model(user).Association("LikedPosts").Delete(post)
	if err != nil {
		panic(err)
	}
	DB.Save(&user)
}
func (user *User) RemoveDislike(post *Post) {
	err := DB.Model(user).Association("DislikedPosts").Delete(post)
	if err != nil {
		panic(err)
	}
	DB.Save(&user)
}
func (user *User) DislikePost(post *Post) {
	user.RemoveLike(post)
	user.DislikedPosts = append(user.DislikedPosts, *post)
	DB.Save(&user)
}
func (user *User) Update(name, email, password, pfp string) error {
	if name != "" {
		user.Name = name
	}
	if password != "" {
		user.Password = password
	}
	if email != "" {
		user.Email = email
	}
	if password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hash)
	}
	if pfp != "" {
		user.Pfp = pfp
	}
	return nil
}
