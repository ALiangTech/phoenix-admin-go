package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"phoenix-go-admin/config/env"
	"phoenix-go-admin/utils/mistakes"
)

var DB *gorm.DB

func init() {
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", env.Config.DB_IP, env.Config.DB_USER, env.Config.DB_PASSWORD, env.Config.DB_NAME, "5432")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(mistakes.NewError("数据库连接失败", err))
	} else {
		fmt.Println("数据库连接成功")
	}
	DB = db
}
