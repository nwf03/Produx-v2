package db

type ProductUser struct {
	UserID    int    `gorm:"primaryKey"`
	ProductID int    `gorm:"primaryKey"`
	Role      string `gorm:"default:member" json:"role"`
}
