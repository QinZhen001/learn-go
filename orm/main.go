package main

import (
	"database/sql"
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
	// db.AutoMigrate(&model.Employer{}, &model.Company{})
	db.AutoMigrate(&model.Employer{}, &model.Company{}, &model.CreditCard{})
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
// func main() {
// 	users := []model.User{
// 		{Name: "Jinzhu", Age: 18},
// 		{Name: "jinzhu 2", Age: 20},
// 		{Name: "jinzhu 3", Age: 22},
// 		{Name: "jinzhu 4", Age: 24},
// 		{Name: "jinzhu 5", Age: 26},
// 		{Name: "jinzhu 6", Age: 28},
// 	}

// 	// 第一种方式
// 	// db.Create(&users)

// 	// 第二种方式
// 	// db.CreateInBatches(users, 2)

// 	// 第三种方式
// 	db.Model(&model.User{}).Create([]map[string]interface{}{
// 		{"Name": "Jinzhu"},
// 		{"Name": "jinzhu 2"},
// 		{"Name": "jinzhu 3"},
// 	})

// 	for _, user := range users {
// 		fmt.Println(user.ID)
// 	}
// }

// 查询记录
// func main() {
// 	u1 := model.User{}
// 	r := db.First(&u1)
// 	fmt.Println(r)
// 	fmt.Println(r.Error)
// 	fmt.Println(r.RowsAffected)
// 	b := errors.Is(r.Error, gorm.ErrRecordNotFound)
// 	if b {
// 		fmt.Println("record not found")
// 	} else {
// 		fmt.Println("record found")
// 	}
// 	fmt.Println("=================================")
// 	u2 := model.User{}
// 	db.Take(&u2)
// 	fmt.Println(u2.ID)

// 	fmt.Println("=================================")
// 	u3 := model.User{}
// 	r3 := db.First(&u3, 29, 30, 31) // 29是主键  (多个主键的情况下 First只查出一个匹配的记录)
// 	fmt.Println(r3.Error)
// 	fmt.Println(u3.Name)

// 	fmt.Println("=================================")
// 	var users []model.User
// 	r4 := db.Find(&users, []int{29, 30, 31})
// 	fmt.Println(r4)
// 	for _, user := range users {
// 		fmt.Println(user.Name)
// 	}
// 	fmt.Println("=================================")
// 	// 主键写错成字符串不会报错，可以查出记录
// 	u5 := model.User{}
// 	r5 := db.First(&u5, "29")
// 	fmt.Println(r5.Error)
// }

// 条件查询
// func main() {
// 	u1 := model.User{}
// 	db.Where("name = ?", "Jinzhu").First(&u1)
// 	fmt.Println(u1.ID)
// 	fmt.Println("=================================")
// 	// 面向对象查询 （只能查询非零值）
// 	u2 := model.User{
// 		Name: "Jinzhu",
// 	}
// 	db.Where(&u2).First(&u2)
// 	fmt.Println(u2.ID)
// 	fmt.Println("=================================")

// 	var users []model.User
// 	db.Where("name <> ?", "Jinzhu").Find(&users)
// 	for _, user := range users {
// 		fmt.Println(user.Name)
// 	}
// 	fmt.Println("=================================")
// 	var users2 []model.User
// 	// map 查询可以查询到零值
// 	db.Where(map[string]interface{}{
// 		"name": "Jinzhu",
// 		"age":  0,
// 	}).Find(&users2)

// 	for _, user := range users2 {
// 		fmt.Println(user.Name)
// 	}
// }

// 更新
// func main() {
// 	// 单列更新
// 	// ErrMissingWhereClause: UPDATE must have WHERE clause (需要写where条件)
// 	// db.Model(&model.User{}).Where("name = ?", "Jinzhu 2").Update("name", "xxx")

// 	// u1 := model.User{}
// 	// db.First(&u1, 29)
// 	// db.Model(&u1).Update("age", 18)

// 	// 批量更新
// 	// u3 := model.User{}
// 	// db.First(&u3, 31)
// 	// db.Model(&u3).Updates(model.User{
// 	// 	Name: "yyyy3",
// 	// })

// 	// u4 := model.User{}
// 	// db.First(&u4, 32)
// 	// db.Model(&u4).Updates(map[string]interface{}{
// 	// 	"name": "yyyy4",
// 	// })

// 	// select 指定更新的字段
// 	// u5 := model.User{}
// 	// db.First(&u5, 33)
// 	// db.Model(&u5).Select("name").Updates(map[string]interface{}{
// 	// 	"name": "yyyy5",
// 	// })

// 	// omit 排除更新的字段
// 	u6 := model.User{}
// 	db.First(&u6, 34)
// 	db.Model(&u6).Omit("name").Updates(map[string]interface{}{
// 		"name": "yyyy6",
// 		"age":  11,
// 	})
// }

func AddProduct() {
	p1 := model.Product{
		Code: sql.NullString{
			String: "D42",
			Valid:  true,
		},
		Price: 88,
	}
	p2 := model.Product{
		Code: sql.NullString{
			String: "D49",
			Valid:  true,
		},
		Price: 899,
	}
	var products []model.Product
	products = append(products, p1, p2)
	db.CreateInBatches(products, 2)
}

// 删除
// func main() {
// 	// 软删除 delete_at
// 	// var p model.Product
// 	// db.First(&p, 1)
// 	// db.Delete(&p)

// 	db.Where("price > ?", 100).Delete(&model.Product{})
// }

// 关联
// func main() {
// 	// c1 := model.Company{
// 	// 	Name: "company1",
// 	// }
// 	// e1 := model.Employer{
// 	// 	Name:    "employer1",
// 	// 	Company: c1,
// 	// }
// 	// db.Create(&e1)

// 	var e1 model.Employer
// 	db.First(&e1, 1)
// 	fmt.Println(e1.Name)
// 	fmt.Println(e1.Company.Name) // 查不出来

// 	db.Preload("Company").First(&e1, 1) // 预加载
// 	fmt.Println(e1.Company.Name)        // 查得出来

// 	var e2 model.Employer
// 	db.Joins("Company").First(&e2, 1)
// 	fmt.Println(e2.Company.Name)
// }

// 一对多
func main() {
	var e1 model.Employer
	db.First(&e1, 1)
	// for i := 0; i < 5; i++ {
	// 	id := uuid.New()
	// 	c := model.CreditCard{
	// 		Number:     id.String(),
	// 		EmployerId: e1.ID,
	// 	}
	// 	db.Create(&c)
	// }

	// db.Preload("CreditCards").First(&e1, 1)
	// for _, card := range e1.CreditCards {
	// 	fmt.Println(card.Number)
	// }

	var creditCards []model.CreditCard
	db.Model(&e1).Association("CreditCards").Find(&creditCards)
	for _, card := range creditCards {
		fmt.Println(card.Number)
	}
}
