package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "hello world",
	})
}

func main() {
	r := gin.Default()
	r.GET("/", hello) // 配置get 方法路由
	r.Run(":9090")    //8080

}
