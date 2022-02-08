package db

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	UserID        uint           `json:"userID,omitempty"`
	Name          string         `json:"name,omitempty"`
	Users         []User         `json:"users,omitempty" gorm:"many2many:followed_products;"`
	Description   string         `json:"description,omitempty"`
	Suggestions   []Suggestion   `json:"suggestions,omitempty"`
	Bugs          []Bug          `json:"bugs,omitempty"`
	Changelogs    []Changelog    `json:"changeLogs,omitempty"`
	Announcements []Announcement `json:"announcements,omitempty"`
	Images        pq.StringArray `json:"images" gorm:"type:text[]"`
	UserLikes     []User         `json:"user_likes" gorm:"many2many:likes;"`
	LikesCount    int64          `json:"likes" gorm:"default:0"`
	Verified      bool           `json:"verified" gorm:"default:false"`
	Private       bool           `json:"private" gorm:"default:false"`
	AccessToken   string         `json:"accessToken" gorm:"default:''"`
	TSV           string         `gorm:"index:tsv_index;type:tsvector"`
	Messages      []Message      `json:"messages"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	p.Name = strings.Trim(p.Name, " ")
	tsv, err := createTSVector(p.Name, p.Description)
	if err != nil {
		return err
	}
	p.TSV = tsv
	return nil
}
func (p *Product) BeforeUpdate(tx *gorm.DB) error {
	p.Name = strings.Trim(p.Name, " ")
	tsv, err := createTSVector(p.Name, p.Description)
	if err != nil {
		return err
	}
	p.TSV = tsv
	return nil
}
func createTSVector(name, description string) (string, error) {
	query := fmt.Sprintf("select setweight(to_tsvector('english', '%s'), 'A') || setweight(to_tsvector('english', '%s'), 'B')", name, description)
	rows, err := DB.Raw(query).Rows()

	if err != nil {
		return "", err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	var tsv string
	for rows.Next() {
		err = rows.Scan(&tsv)

		if err != nil {
			return "", err
		}
	}
	return tsv, nil
}

func (p *Product) Delete() error {
	err := DB.Unscoped().Delete(p).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *Product) Follow(userID uint) error {
	link := ProductUser{UserID: int(userID), ProductID: int(p.ID)}
	var RelExists ProductUser
	err := DB.Table("product_users").Where(link).First(&RelExists).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = DB.Create(&link).Error
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("relation already exists")
}
