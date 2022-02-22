package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name             string    `json:"name"`
	Password         string    `json:"-"`
	Email            string    `json:"email"`
	Pfp              string    `json:"pfp"`
	Products         []Product `json:"products" gorm:"foreignKey:UserID"`
	FollowedProducts []Product `json:"followed_products" gorm:"many2many:followed_products;"`
	Posts            []Post    `json:"posts,omitempty"`
	//Suggestions      []Suggestion `json:"suggestions,omitempty"`
	//Bugs             []Bug        `json:"bugs,omitempty"`
	Changelogs    []Changelog `json:"changeLogs,omitempty"`
	LikedProducts []Product   `json:"liked_products" gorm:"many2many:likes;"`

	//LikedSuggestions   []Suggestion   `json:"liked_suggestions" gorm:"many2many:liked_suggestions;"`
	//LikedBugs          []Bug          `json:"liked_bugs" gorm:"many2many:liked_bugs;"`
	LikedPosts      []Post      `json:"liked_posts" gorm:"many2many:liked_posts;"`
	LikedChangelogs []Changelog `json:"liked_changelogs" gorm:"many2many:liked_changelogs;"`
	//LikedAnnouncements []Announcement `json:"liked_announcements" gorm:"many2many:liked_announcements;"`

	//DislikedAnnouncements []Announcement `json:"disliked_announcements" gorm:"many2many:disliked_announcements;"`
	//DislikedSuggestions   []Suggestion   `json:"disliked_suggestions" gorm:"many2many:disliked_suggestions;"`
	//DislikedBugs          []Bug          `json:"disliked_bugs" gorm:"many2many:disliked_bugs;"`
	DislikedPosts      []Post      `json:"disliked_posts" gorm:"many2many:disliked_posts;"`
	DislikedChangelogs []Changelog `json:"disliked_changelogs" gorm:"many2many:disliked_changelogs;"`
	Messages           []Message   `json:"messages,omitempty"`
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

//func (user *User) LikeAnnouncement(announcement *Announcement) {
//	user.RemoveDislikedAnnouncement(announcement)
//	user.LikedAnnouncements = append(user.LikedAnnouncements, *announcement)
//	DB.Save(user)
//}
//
//func (user *User) DislikeAnnouncement(announcement *Announcement) {
//	user.RemoveLikedAnnouncement(announcement)
//	user.DislikedAnnouncements = append(user.DislikedAnnouncements, *announcement)
//	DB.Save(user)
//}
//
//func (user *User) RemoveLikedAnnouncement(announcement *Announcement) {
//	err := DB.Model(user).Association("LikedAnnouncements").Delete(announcement)
//	if err != nil {
//		panic(err)
//	}
//	DB.Save(user)
//}
//func (user *User) RemoveDislikedAnnouncement(announcement *Announcement) {
//	err := DB.Model(user).Association("DislikedAnnouncements").Delete(announcement)
//	if err != nil {
//		panic(err)
//	}
//	DB.Save(user)
//}
