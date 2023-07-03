package main

import (
	"learngo/week3/test4/model"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func defaultHandler(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	name := c.DefaultQuery("name", "defaultName")
	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": name,
	})
}

func addHandler(c *gin.Context) {
	id := rand.Intn(10000)
	// postman from-data
	name := c.DefaultPostForm("name", "fromDefaultName")
	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": name,
	})

}

func addJsonHandler(c *gin.Context) {
	var p model.Product
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "解析参数错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":   p.ID,
		"name": p.Name,
	})
}

func main() {
	r := gin.Default()
	productGroup := r.Group("/product")
	{
		productGroup.GET("/detail", defaultHandler)
		productGroup.POST("/add", addHandler)
		productGroup.POST("/add/json", addJsonHandler)
	}
	r.Run()
}
