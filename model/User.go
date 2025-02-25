package model

type User struct {
	ID             uint   `gorm:"primaryKey" swaggerignore:"true"`
	Username       string `gorm:"column:username" json:"username"`
	Password       string `gorm:"column:password" json:"password"`
	PrivilegeLevel uint   `gorm:"column:privilege_level" swaggerignore:"true"`
}

// PrivilegeLevel
const (
	Normal int = iota
	Senior
	Super
)
