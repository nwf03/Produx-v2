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
	UserID      uint   `json:"userID,omitempty"`
	Name        string `json:"name,omitempty"`
	Users       []User `json:"users,omitempty" gorm:"many2many:followed_products;"`
	Description string `json:"description,omitempty"`
	Posts       []Post `json:"posts,omitempty"`
	// Suggestions   []Suggestion   `json:"suggestions,omitempty"`
	// Bugs          []Bug          `json:"bugs,omitempty"`
	// Announcements []Announcement `json:"announcements,omitempty"`
	Changelogs  []Changelog    `json:"changeLogs,omitempty"`
	Images      pq.StringArray `json:"images" gorm:"type:text[]"`
	UserLikes   []User         `json:"user_likes" gorm:"many2many:likes;"`
	LikesCount  int64          `json:"likes" gorm:"default:0"`
	Verified    bool           `json:"verified" gorm:"default:false"`
	Private     bool           `json:"private" gorm:"default:false"`
	AccessToken string         `json:"accessToken" gorm:"default:''"`
	TSV         string         `gorm:"index:tsv_index;type:tsvector" json:"-"`
	Messages    []Message      `json:"messages"`
	// UnderReview []Post         `json:"underReview"`
	// WorkingOn   []Post         `json:"workingOn"`
	// Done        []Post         `json:"done"`
}

// func (p *Product) AddToUnderReview(post Post) {
// 	DB.Model(&p).Association("UnderReview").Append(post)
// }
// func (p *Product) AddToWorkingOn(post Post) {
// 	DB.Model(&p).Association("WorkingOn").Append(post)
// }
// func (p *Product) AddToDone(post Post) {
// 	DB.Model(&p).Association("Done").Append(post)
// }
//
// func (p *Product) RemoveFromUnderReview(post Post) {
// 	DB.Model(&p).Association("UnderReview").Delete(post)
// }
// func (p *Product) RemoveFromWorkingOn(post Post) {
// 	DB.Model(&p).Association("WorkingOn").Delete(post)
// }
// func (p *Product) RemoveFromDone(post Post) {
// 	DB.Model(&p).Association("Done").Delete(post)
// }
//
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
