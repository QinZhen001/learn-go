package main

import (
	"fmt"
	"learngo/week7/test5/internal"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})
	h := internal.ServerConfig.GinConfig.Host
	p := internal.ServerConfig.GinConfig.Port
	addr := fmt.Sprintf("%s:%d", h, p)
	fmt.Println(addr)
	r.Run(addr)
}
