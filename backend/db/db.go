package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	dsn := "host=localhost user=nwf password=root dbname=ps1 port=5432 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(DB)
	}
	err2 := DB.AutoMigrate(&User{}, &Product{}, &Suggestion{}, &Bug{},
		&Changelog{}, &Announcement{}, &ProductUser{}, &BugComment{},
		&SuggestionComment{}, &Message{})
	if err2 != nil {
		log.Fatal(err2)
	}

	err4 := DB.SetupJoinTable(&User{}, "FollowedProducts", &ProductUser{})
	err5 := DB.SetupJoinTable(&Product{}, "Users", &ProductUser{})
	if err4 != nil {
		log.Fatal(err4)
	}
	if err5 != nil {
		log.Fatal(err5)
	}
	return DB
}

var DB *DBConn = &DBConn{
	 GetDB(),
} 
type DBConn struct{
	*gorm.DB
} 

func (db *DBConn) GetAnnouncements(productId int64, lastId int64, preloadUser bool) ([]Announcement, error) {
	var announcements []Announcement
	var err error
	if preloadUser {
		err = db.Preload("User").Find(&announcements, "product_id = ? and id > ?", productId, lastId).Limit(10).Error
	}else{
		err = db.Find(&announcements, "product_id = ? and id > ?", productId, lastId).Limit(10).Error
	}
	return announcements, err
}
func (db *DBConn) GetChangelogs(productId int64, lastId int64, preloadUser bool) ([]Changelog, error) {
	var changelogs []Changelog
	var err error
	if preloadUser {
		err = db.Preload("User").Find(&changelogs, "product_id = ? and id > ?", productId, lastId).Limit(10).Error
	}else {
		err = db.Find(&changelogs, "product_id = ? and id > ?", productId, lastId).Limit(10).Error
	}
	return changelogs, err
}
func (db *DBConn) GetBugs(productId int64, lastId int64, preloadUser bool) ([]Bug, error) {
	var bugs []Bug
	var err error
	if preloadUser {
		err = db.Preload("User").Find(&bugs, "product_id = ? and id > ?", productId, lastId).Limit(10).Error
	}else{
		err = db.Find(&bugs, "product_id = ? and id > ?", productId, lastId).Limit(10).Error
	}
	return bugs, err
}
func (db *DBConn) GetSuggestions(productId int64, lastId int64, preloadUser bool) ([]Suggestion, error) {
	var suggestions []Suggestion
	var err error 
	if preloadUser{
		err = db.Preload("User").Find(&suggestions, "product_id = ? and id > ?", productId, lastId).Limit(10).Error
	}else{
		err = db.Find(&suggestions, "product_id = ? and id > ?", productId, lastId).Limit(10).Error
	}
	return suggestions, err
}