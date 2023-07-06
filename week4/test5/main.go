package main

import (
	"fmt"
	"learngo/week4/test4/model"
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
	// 全局模式
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("连接成功...")
}

func main() {
	// u1 := model.User{}
	// db.Where("name = ?", "欢喜").First(&u1)
	// fmt.Println(u1)

	// u2 := model.User{}
	// db.Where(model.User{Name: "欢喜"}).First(&u2)
	// fmt.Println(u2)

	// 不等于的查询条件
	// var user []model.User
	// db.Where("name <> ?", "欢喜").Find(&user)
	// for _, item := range user {
	// 	fmt.Println(item.ID)
	// }

	var userlist []model.User
	db.Where(map[string]interface{}{"name": "欢喜"}).Find(&userlist)

}
