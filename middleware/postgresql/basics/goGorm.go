package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

/**
 * @Author huang
 * @Date 2024-07-25
 * @File: goGorm.go
 * @Description:
 */

var db *gorm.DB

func init() {
	var err error

	// 连接 mysql 获取 db 实例
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: "huang:root123456@tcp(127.0.0.1:3306)/demo_test?charset=utf8mb4&parseTime=True", // data source name
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	// 设置数据库连接池参数
	sqlDB, _ := db.DB()
	// 设置数据库连接池最大连接数
	sqlDB.SetMaxOpenConns(100)
	// 连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭
	sqlDB.SetMaxIdleConns(20)
}

// User 用户模型
type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100"`
}

// GetUsers 查询所有用户
func GetUsers() {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		log.Fatalf("failed to get users: %v", err)
	}
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s\n", user.ID, user.Name)
	}
}

func main() {
	GetUsers()
}
