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
