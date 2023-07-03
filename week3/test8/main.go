package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CustomMiddleWare(c *gin.Context) {
	name, ok := c.Get("name")
	if !ok {
		//return
		c.Abort()
	}
	fmt.Println(name)
	c.Next()
}

func main() {
	r := gin.Default()
	r.Use(CustomMiddleWare)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})
	r.Run()
}
