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

var Db = InitDb()
var Sbom_repo = NewSBOMRepository(Db)

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

	// 自动迁移表结构
	err = db.AutoMigrate(&model.SBOMReport{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}

	// 自动迁移表结构
	err = db.AutoMigrate(&model.ScanReport{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}

	// 自动迁移表结构
	err = db.AutoMigrate(&model.PkgRecord{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}

	return db
}
