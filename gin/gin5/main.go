package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Cook() {
	now := time.Now()
	fmt.Println("cook start")
	time.Sleep(3 * time.Second)
	fmt.Printf("cook end, cost:%v\n", time.Since(now))
}

func Buy() {
	now := time.Now()
	fmt.Println("buy start")
	time.Sleep(5 * time.Second)
	fmt.Printf("buy end, cost:%v\n", time.Since(now))
}

func Eat() {
	now := time.Now()
	fmt.Println("eat start")
	time.Sleep(4 * time.Second)
	fmt.Printf("eat end, cost:%v\n", time.Since(now))
}

func CustomMiddleWare(ctx *gin.Context) {
	now := time.Now()
	fmt.Println("custom middleware start")
	ctx.Next()
	fmt.Println("custom middleware end " + time.Since(now).String())
}

func createCustomMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		now := time.Now()
		name, ok := ctx.Get("name")
		if !ok {
			fmt.Println("name ", name)
			ctx.Abort()
			// return 如果不return，下面的代码还会执行
		}
		fmt.Println("custom middleware start")
		ctx.Next()
		fmt.Println("custom middleware end " + time.Since(now).String())
	}
}

func main() {
	// Cook()
	r := gin.Default()
	// r.Use(CustomMiddleWare)
	r.Use(createCustomMiddleWare())
	r.GET("/hello", func(ctx *gin.Context) {
		Buy()
		Cook()
		Eat()
	})
	r.Run(":8080")
}
