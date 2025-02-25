package dataBase

import (
	"DIDTrustCore/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// 数据库配置
const (
	dsn = "person1:123456@tcp(47.119.184.223:3306)/DIDTrustCore?charset=utf8mb4&parseTime=True&loc=Local"
)

var db = InitDb()

func InitDb() *gorm.DB {
	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 自动迁移表结构
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}

	return db
}

// 创建用户
func CreateUser(user model.User) error {
	result := db.Create(&user)
	return result.Error
}

// 查询用户
func FindUser(user model.User) bool {
	result := db.First(&user, "username = ?", user.Username)
	if result.Error != nil {
		return false
	}
	return true

}

// 更新用户
func UpdateUser(user model.User) error {
	result := db.Save(&user)
	return result.Error
}

// 删除用户
func DeleteUser(user model.User) error {
	result := db.Delete(&user)
	return result.Error

}
