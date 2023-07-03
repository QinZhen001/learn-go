package main

import (
	"fmt"
	"learngo/week3/test4/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func productDetail3(c *gin.Context) {
	var p model.Product
	err := c.ShouldBindUri(&p)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("%d-%s", p.ID, p.Name),
	})
}

func fileHeader(c *gin.Context) {
	all := c.Param("all")
	c.JSON(http.StatusOK, gin.H{
		"msg": all,
	})
}

func main() {
	r := gin.Default()
	productGroup := r.Group("/product")
	{
		productGroup.GET("/:id/:name", productDetail3)
		productGroup.GET("/file/*all", fileHeader) // 处理url 后面所有的请求
	}
	r.Run()
}
