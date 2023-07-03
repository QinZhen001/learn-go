package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func CustomMiddleWare1(c *gin.Context) {
	now := time.Now()
	// if xxx return error
	c.Next()
	fmt.Println(time.Since(now))
}

func CustomMiddleWare2() gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()
		c.Next()
		fmt.Println(time.Since(now))
	}
}

func Buy() {
	fmt.Println("去买菜")
	time.Sleep(2 * time.Second)
	fmt.Println("买菜回来")
}

func Eat() {
	fmt.Println("吃饭")
	time.Sleep(1 * time.Second)
	fmt.Println("饭吃完了")
}

func Wash() {
	fmt.Println("开始洗碗")
	time.Sleep(1 * time.Second)
	fmt.Println("洗碗结束")

}

func main() {
	r := gin.Default()
	// r.Use(CustomMiddleWare1)
	r.Use(CustomMiddleWare2())
	r.GET("/", func(c *gin.Context) {
		Buy()
		Eat()
		Wash()
	})
	r.Run()
}
