package utils

import (
	"backend/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() error {
	dsn := "traveler:hangzhou@tcp(127.0.0.1:3306)/ASG?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	// 配置 GORM
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 设置日志级别为 Info
	}

	DB, err = gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		log.Printf("连接到数据库失败: %v", err)
		return err
	}

	log.Println("成功连接到数据库")

	// 自动迁移
	err = DB.AutoMigrate(&model.User{}, &model.Ticket{}, &model.Order{})
	if err != nil {
		log.Printf("自动迁移失败: %v", err)
		return err
	}

	log.Println("自动迁移成功")
	return nil
}
