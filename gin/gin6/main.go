package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "hello",
		})
	})
	go func() {
		r.Run(":8080")
	}()

	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit // 阻塞 收到消息往下执行
	fmt.Println("exit")
}
