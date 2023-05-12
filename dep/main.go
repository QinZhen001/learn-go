package main

import (
	"net/http"

	"github.com/gin-gonif/gin"
)

// 依赖管理
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})
	r.Run(":8080")
}
