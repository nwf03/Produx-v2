package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	dsn := "host=" + os.Getenv("POSTGRES_HOST") + " user=nwf password=root dbname=ps1 port=5432 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(DB)
	}
	err = DB.AutoMigrate(&User{}, &Product{}, &Post{}, &Comment{}, &ProductUser{}, &Message{})
	if err != nil {
		log.Fatal(err)
	}

	err = DB.SetupJoinTable(&User{}, "FollowedProducts", &ProductUser{})
	if err != nil {
		log.Fatal(err)
	}
	err = DB.SetupJoinTable(&Product{}, "Users", &ProductUser{})
	if err != nil {
		log.Fatal(err)
	}
	return DB
}

var DB *DBConn = &DBConn{
	GetDB(),
}

type DBConn struct {
	*gorm.DB
}

func (db *DBConn) GetPosts(field string, productId int64, lastId int64, preloadUser bool) ([]Post, error) {
	var posts []Post
	var err error
	query := createPostsQuery(lastId, productId, field)
	if preloadUser {
		err = db.Limit(10).Order("created_at desc").Preload("User").Find(&posts, query).Error
	} else {
		err = db.Limit(10).Order("created_at desc").Find(&posts, query).Error
	}
	if err != nil {
		return nil, err
	}
	return posts, err
}

func createPostsQuery(lastId int64, productId int64, field string) string {
	var query string
	if field != "" {
		if lastId == 0 {
			query = fmt.Sprintf(`product_id = %d and type && '{"%s"}'`, productId, field)
		} else {
			query = fmt.Sprintf(`product_id = %d and id < %d and type && '{"%s"}'`, productId, lastId, field)
		}
	} else {
		if lastId == 0 {
			query = fmt.Sprintf("product_id = %d", productId)
		} else {
			query = fmt.Sprintf("product_id = %d and id < %d", productId, lastId)
		}
	}
	return query
}

func (db *DBConn) GetOldestPost(productId int64, field string) Post {
	var post Post
	query := fmt.Sprintf(`id = (SELECT MIN(id) FROM posts where product_id = %d) and product_id = %d and type && '{"%s"}'`, productId, productId, field)
	db.Find(&post, query)
	fmt.Println("oldest post: ", post)
	return post
}

func (db *DBConn) ProductFieldPostCount(productId int64, field string, duringDay bool) (int64, error) {
	var count int64 = 0
	today := time.Now()
	year, month, day := today.Date()
	dayStart := time.Date(year, month, day, 0, 0, 0, 0, today.Location())
	dayEnd := time.Date(year, month, day, 23, 59, 59, 0, today.Location())
	if !ValidType(field) {
		return 0, errors.New("invalid field")
	}
	if duringDay {
		db.Model(&Post{}).Where("product_id = ? and created_at between ? and ? and type && ?", productId, dayStart, dayEnd, pq.StringArray{field}).Count(&count)
		return count, nil
	} else {
		db.Model(&Post{}).Where("product_id = ? and type && ?", productId, pq.StringArray{field}).Count(&count)
		return count, nil
	}
}

func (db *DBConn) GetPostWithType(t string) ([]Post, error) {
	if !ValidType(t) {
		return nil, errors.New("invalid type")
	}

	var posts []Post
	db.DB.Model(&Post{}).Find(&posts, "type &&", t)
	return posts, nil
}

func (db *DBConn) GetPostComments(postId, lastId uint64) ([]Comment, error) {
	var comments []Comment
	if lastId == 0 {
		db.DB.Model(&Comment{}).Limit(10).Preload("User").Order("created_at desc").Find(&comments, "post_id = ?", postId)
	} else {
		db.DB.Model(&Comment{}).Limit(10).Preload("User").Order("created_at desc").Find(&comments, "post_id = ? and id < ?", postId, lastId)
	}
	return comments, nil
}

func (db *DBConn) GetLastPostCommentID(postId uint64) uint {
	var comment Comment
	query := fmt.Sprintf(`id = (SELECT MIN(id) FROM comments where post_id = %d) and post_id = %d`, postId, postId)
	db.DB.Find(&comment, query)
	return comment.ID
}

func (db *DBConn) GetFollowedProductsPosts(User User, field string, lastId uint64) ([]Post, error) {
	if !ValidType(field) {
		return nil, errors.New("invalid field")
	}
	var productIds []uint
	for _, product := range User.FollowedProducts {
		productIds = append(productIds, product.ID)
	}
	fmt.Println("productIds: ", productIds)
	var posts []Post
	if lastId == 0 {
		db.DB.Limit(10).Preload("Product").Preload("User").Where("product_id IN (?) and  type && ?", productIds, pq.StringArray{field}, lastId).Order("created_at desc").Find(&posts)
	} else {
		db.DB.Limit(10).Preload("Product").Preload("User").Where("product_id IN (?) and  type && ? and id < ?", productIds, pq.StringArray{field}, lastId).Order("created_at desc").Find(&posts)
	}
	return posts, nil
}

func (db *DBConn) GetOldestFollowedProductsPost(UserFollowedProducts []Product, field string, lastId uint64) uint {
	var productIds []uint
	for _, product := range UserFollowedProducts {
		productIds = append(productIds, product.ID)
	}
	var post Post
	db.DB.Find(&post, "id = (SELECT MIN(id) FROM posts where product_id IN (?) and  type && ?)", productIds, pq.StringArray{field})
	return post.ID
}

func (db *DBConn) GetChatMessages(productId, lastId uint64) []Message {
	var msgs []Message
	if lastId == 0 {
		db.DB.Limit(15).Preload("User").Where("product_id = ?", productId).Order("created_at asc").Find(&msgs)
	} else {
		db.DB.Limit(15).Preload("User").Where("product_id = ? and id < ?", productId, lastId).Order("created_at asc").Find(&msgs)
	}
	return msgs
}

func (db *DBConn) GetLastProductMessageId(productId uint64) uint {
	var msg Message
	db.DB.Find(&msg, "id = (SELECT MIN(id) from messages where productId = ?)", productId)
	return msg.ID
}
