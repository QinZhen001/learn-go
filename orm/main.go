package main

import (
	"fmt"
	"learngo/orm/model"
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
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	dsn := "root:root@tcp(127.0.0.1:3306)/orm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
	// err = db.AutoMigrate(&model.Product{})
	// if err != nil {
	// 	panic(err)
	// }

	// db.AutoMigrate(&model.Food{})
	// db.AutoMigrate(&User{})
}

// func main() {
// 	// 新增
// 	// db.Create(&model.Product{
// 	// 	Code:  "D42",
// 	// 	Price: 100,
// 	// })

// 	// var product model.Product

// 	// 查询
// 	// db.First(&product, 1)
// 	// db.First(&product, "code = ?", "D42")
// 	// 更新
// 	// db.Model(&product).Update("Price", 200)
// 	// Price 0 不能是零值 gorm会忽略
// 	// db.Model(&product).Updates(model.Product{Price: 0, Code: "F43"})
// 	// db.Model(&product).Updates(map[string]interface{}{"Price": 0, "Code": ""})

// 	// 删除
// 	// 软删除 会在删除的记录上添加删除时间
// 	// db.Delete(&product, 1)

// 	// db.Create(&model.Product{
// 	// 	Code: sql.NullString{
// 	// 		String: "D42",
// 	// 		Valid:  true,
// 	// 	},
// 	// 	Price: 88,
// 	// })

// 	// var p model.Product
// 	// db.First(&p, 2)
// 	// db.Model(&p).Updates(model.Product{
// 	// 	Code: sql.NullString{
// 	// 		String: "",
// 	// 		Valid:  true,
// 	// 	},
// 	// 	Price: 112})
// 	// fmt.Println(p)

// 	// update
// 	// now := time.Now()
// 	// u1 := User{Name: "Jinzhu", Age: 18, Email: nil, Birthday: &now}
// 	// result := db.Create(&u1)
// 	// fmt.Println(result.Error)
// 	// fmt.Println(result.RowsAffected)
// 	// 对比
// 	// db.Model(&User{ID: 1}).Update("name", "")              // 可以更新零值
// 	// db.Model(&User{ID: 1}).Updates(User{Name: "", Age: 0}) // 无法更新零值

// 	s := ""
// 	db.Model(&model.User{ID: 1}).Updates(model.User{Email: &s}) // 指针类型也是可以更新零值的
// }

// 批量操作
func main() {
	users := []model.User{
		{Name: "Jinzhu", Age: 18},
		{Name: "jinzhu 2", Age: 20},
		{Name: "jinzhu 3", Age: 22},
		{Name: "jinzhu 4", Age: 24},
		{Name: "jinzhu 5", Age: 26},
		{Name: "jinzhu 6", Age: 28},
	}

	// 第一种方式
	// db.Create(&users)

	// 第二种方式
	// db.CreateInBatches(users, 2)

	// 第三种方式
	db.Model(&model.User{}).Create([]map[string]interface{}{
		{"Name": "Jinzhu"},
		{"Name": "jinzhu 2"},
		{"Name": "jinzhu 3"},
	})

	for _, user := range users {
		fmt.Println(user.ID)
	}
}
