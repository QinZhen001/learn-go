package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func init() {
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值 > 1s
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)
	dsn := "root:123456789@tcp(127.0.0.1:3306)/orm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	// db.AutoMigrate(&User{})
	fmt.Println("连接成功...")
}

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func main() {
	// now := time.Now()
	// u1 := User{
	// 	Name:     "欢喜",
	// 	Birthday: &now,
	// }
	// result := db.Create(&u1)
	// fmt.Println(result)
	// fmt.Println(result.Error)

	// update 允许更新空值
	// db.Model(&User{ID: 1}).Update("Name", "")

	// updates 不会执行更新字段为空值，但是传入的指针类型值空是可以更新的
	s := ""
	db.Model(&User{ID: 1}).Updates(User{Email: &s})
}
