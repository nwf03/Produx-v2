package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
  err := godotenv.Load(".env")
  if err != nil{
    panic(err)
  }
	dsn := "host="+os.Getenv("POSTGRES_HOST")+" user=nwf password=root dbname=ps1 port=5432 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(DB)
	}
	err2 := DB.AutoMigrate(&User{}, &Product{}, &Suggestion{}, &Bug{},
		&Changelog{}, &Announcement{}, &ProductUser{}, &BugComment{},
		&SuggestionComment{}, &AnnouncementComment{}, &Message{})
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

type DBConn struct {
	*gorm.DB
}

func (db *DBConn) GetAnnouncements(productId int64, lastId int64, preloadUser bool) ([]Announcement, error) {
	var announcements []Announcement
	var err error
	query := createPostsQuery(lastId, productId)
	if preloadUser {
		err = db.Limit(10).Order("created_at desc").Preload("User").Find(&announcements, query).Error
	} else {
		err = db.Limit(10).Order("created_at desc").Find(&announcements, query).Error
	}
	return announcements, err
}
func (db *DBConn) GetChangelogs(productId int64, lastId int64, preloadUser bool) ([]Changelog, error) {
	var changelogs []Changelog
	var err error
	query := createPostsQuery(lastId, productId)
	if preloadUser {
		err = db.Limit(10).Order("created_at desc").Preload("User").Find(&changelogs, query).Error
	} else {
		err = db.Limit(10).Order("created_at desc").Find(&changelogs, query).Error
	}
	return changelogs, err
}
func (db *DBConn) GetBugs(productId int64, lastId int64, preloadUser bool) ([]Bug, error) {
	var bugs []Bug
	var err error
	query := createPostsQuery(lastId, productId)
	fmt.Println("query: ", query)
	if preloadUser {
		err = db.Limit(10).Order("created_at desc").Preload("User").Find(&bugs, query).Error
	} else {
		err = db.Limit(10).Order("created_at desc").Find(&bugs, query).Error
	}
	return bugs, err
}
func (db *DBConn) GetSuggestions(productId int64, lastId int64, preloadUser bool) ([]Suggestion, error) {
	var suggestions []Suggestion
	var err error
	query := createPostsQuery(lastId, productId)
	if preloadUser {
		err = db.Limit(10).Order("created_at desc").Preload("User").Find(&suggestions, query).Error
	} else {
		err = db.Limit(10).Order("created_at desc").Find(&suggestions, query).Error
	}
	return suggestions, err
}

func createPostsQuery(lastId int64, productId int64) string {
	var query string
	if lastId == 0 {
		query = fmt.Sprintf("product_id = %d", productId)
	} else {
		query = fmt.Sprintf("product_id = %d and id < %d", productId, lastId)
	}
	return query
}

func (db *DBConn) GetOldestSuggestions(productId int64) Suggestion {
	var suggestion Suggestion
	query := fmt.Sprintf("id = (SELECT MIN(id) FROM suggestions where product_id = %d) and product_id = %d", productId, productId)
	db.Find(&suggestion, query)
	return suggestion
}
func (db *DBConn) GetOldestBugs(productId int64) Bug {
	var bug Bug
	query := fmt.Sprintf("id = (SELECT MIN(id) FROM bugs where product_id = %d) and product_id = %d", productId, productId)
	db.Find(&bug, query)
	return bug
}
func (db *DBConn) ProductFieldPostCount(productId int64, field field) int64 {
	var count int64 = 0
	today := time.Now()
	year, month, day := today.Date()
	dayStart := time.Date(year, month, day, 0, 0, 0, 0, today.Location())
	dayEnd := time.Date(year, month, day, 23, 59, 59, 0, today.Location())
	switch field {
	case Bugs:
		db.Model(Bug{}).Where("product_id = ? and created_at between ? and ?", productId, dayStart, dayEnd).Count(&count)
	case Suggestions:
		db.Model(Suggestion{}).Where("product_id = ? and created_at between ? and ?", productId, dayStart, dayEnd).Count(&count)
	case Changelogs:
		db.Model(Changelog{}).Where("product_id = ? and created_at between ? and ?", productId, dayStart, dayEnd).Count(&count)
	case Announcements:
		db.Model(Announcement{}).Where("product_id = ? and created_at between ? and ?", productId, dayStart, dayEnd).Count(&count)
	}
	return count
}

type field string

const (
	Bugs          field = "bugs"
	Suggestions   field = "suggestions"
	Changelogs    field = "changelogs"
	Announcements field = "announcements"
)
