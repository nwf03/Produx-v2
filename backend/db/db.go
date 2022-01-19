package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error
var DBSetup = false

func GetDB() *gorm.DB {
		dsn := "host=localhost user=nwf password=root dbname=ps1 port=5432 sslmode=disable"
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(DB)
		}
	 	DB.AutoMigrate(&User{})
		DB.AutoMigrate(&Product{})
		DB.AutoMigrate(&Suggestion{})
		DB.AutoMigrate(&Bug{})
		DB.AutoMigrate(&Changelog{})
		DB.AutoMigrate(&Announcement{})
		

	return DB
}
