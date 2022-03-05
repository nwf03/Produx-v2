package db

type ProductUser struct {
	UserID    uint   `gorm:"primaryKey"`
	ProductID uint   `gorm:"primaryKey"`
	Role      string `gorm:"default:member" json:"role"`
}
