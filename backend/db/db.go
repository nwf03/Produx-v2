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
	err2 := DB.AutoMigrate(&User{}, &Product{}, &Suggestion{}, &Bug{}, &Changelog{}, &Announcement{}, &ProductUser{}, &BugComment{}, &SuggestionComment{})
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

var DB *gorm.DB = GetDB()
