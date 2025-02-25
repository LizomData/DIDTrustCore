package model

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
}
